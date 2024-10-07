package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	conn *websocket.Conn
	send chan []byte
}

type ChatRoom struct {
	clients   map[*Client]bool
	broadcast chan []byte
	mu        sync.Mutex
}

func newChatRoom() *ChatRoom {
	return &ChatRoom{
		clients:   make(map[*Client]bool),
		broadcast: make(chan []byte),
	}
}

func (cr *ChatRoom) run() {
	for {
		msg := <-cr.broadcast
		cr.mu.Lock()
		for client := range cr.clients {
			select {
			case client.send <- msg:
			default:
				close(client.send)
				delete(cr.clients, client)
			}
		}
		cr.mu.Unlock()
	}
}

func handleConnection(cr *ChatRoom, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Ошибка при обновлении соединения:", err)
		return
	}
	client := &Client{conn: conn, send: make(chan []byte)}
	cr.mu.Lock()
	cr.clients[client] = true
	cr.mu.Unlock()

	go client.writePump()
	client.readPump(cr)
}

func (c *Client) readPump(cr *ChatRoom) {
	defer func() {
		c.conn.Close()
	}()
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			fmt.Println("Ошибка при чтении сообщения:", err)
			break
		}
		cr.broadcast <- msg
	}
}

func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()
	for msg := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			fmt.Println("Ошибка при отправке сообщения:", err)
			break
		}
	}
}

func main() {
	chatRoom := newChatRoom()
	go chatRoom.run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleConnection(chatRoom, w, r)
	})

	fmt.Println("Сервер запущен на порту 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}

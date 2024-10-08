package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gorilla/websocket"
	"golang.org/x/sync/errgroup"
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

func (cr *ChatRoom) closeClients() {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	for client := range cr.clients {
		client.conn.WriteMessage(websocket.TextMessage, []byte("connection closed"))
		client.conn.Close()
		delete(cr.clients, client)
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
		cr.mu.Lock()
		delete(cr.clients, c)
		cr.mu.Unlock()
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
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	chatRoom := newChatRoom()
	go chatRoom.run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleConnection(chatRoom, w, r)
	})

	fmt.Println("Сервер запущен на порту 8000")

	httpServer := &http.Server{
		Addr: ":8000",
	}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return httpServer.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		chatRoom.closeClients()
		return httpServer.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
}

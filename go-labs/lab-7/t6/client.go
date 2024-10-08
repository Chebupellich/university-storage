package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8000/ws", nil)
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		return
	}
	defer conn.Close()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		defer conn.Close()
		for {
			mt, msg, err := conn.ReadMessage()
			if mt == websocket.CloseMessage {
				fmt.Println("Сервер завершил работу")
				return
			}
			if err != nil {
				fmt.Println("Ошибка при чтении сообщения:", err)
				close(sig)
				return
			}
			if len(msg) == 0 {
				continue
			}
			fmt.Printf("\rПолучено сообщение: %s\n", msg)
			fmt.Print("> ")
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		defer conn.Close()
		for {
			fmt.Print("> ")
			if scanner.Scan() {
				msg := scanner.Text()
				err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
				if err != nil {
					fmt.Println("Ошибка при отправке сообщения:", err)
					close(sig)
					return
				}
			}
		}
	}()

	go func() {
		<-sig
		err := conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "client disconnecting"), time.Now().Add(time.Second))
		if err != nil {
			fmt.Println("Ошибка при отправке сообщения:", err)
		}
		conn.Close()
	}()

	<-sig

	fmt.Println("\nЗавершение работы...")
}

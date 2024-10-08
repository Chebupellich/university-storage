package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

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

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
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

	<-sig

	fmt.Println("\nЗавершение работы...")
}

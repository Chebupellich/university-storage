package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Error set connection")
		return
	}
	defer conn.Close()

	for {
		var msg string
		fmt.Print("Write message: ")
		_, err := fmt.Scanln(&msg)
		if err != nil {
			fmt.Println("Incorrect input")
			continue
		}

		if n, err := conn.Write([]byte(msg)); n == 0 || err != nil {
			fmt.Println("Error while send message: ", err)
			return
		}

		fmt.Print("Response: ")
		buff := make([]byte, 1024*16)
		n, err := conn.Read(buff)
		if err != nil {
			break
		}

		fmt.Println(string(buff[0:n]))
	}
}

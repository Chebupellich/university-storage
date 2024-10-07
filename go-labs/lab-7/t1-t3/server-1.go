package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8800")
	if err != nil {
		fmt.Println("Error while try run server: ", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server running on ", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error while try connect: ", err)
			conn.Close()
			continue
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		input := make([]byte, (1024 * 64))
		n, err := conn.Read(input)

		fmt.Println("User connected: ", conn.RemoteAddr())

		if n == 0 {
			fmt.Println("User disconected: ", conn.RemoteAddr())
			break
		}

		if err != nil {
			fmt.Println("Error while read ", conn.RemoteAddr(), " : ", err)
			break
		}

		msg := string(input[0:n])
		fmt.Println("User ", conn.RemoteAddr(), " message: ", msg)

		conn.Write([]byte(msg))
	}
}

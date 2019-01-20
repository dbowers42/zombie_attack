package server

import (
	"fmt"
	"log"
	"net"
)

func Run(port int) (err error) {
	host := fmt.Sprintf(":%d", port)

	listener, err := net.Listen("tcp", host)

	if err != nil {
		return
	}

	log.Printf("Server running on port: %d\n", port)

	for {
		conn, err := listener.Accept()

		if err != nil {
			break
		}

		err = handleRequest(conn)
	}

	return
}

func handleRequest(conn net.Conn) error {
	buff := make([]byte, 1024)

	l, err := conn.Read(buff)

	if err != nil {
		return err
	}

	_, err = conn.Write(buff[:l])

	if err != nil {
		return err
	}

	return conn.Close()
}

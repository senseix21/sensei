package main

import (
	"net"
	"sensei/internal"
)

func main() {
	//Create new server
	server := internal.NewServer(":8080")

	// Register routes
	server.Router.Handle("GET", "/", func(conn net.Conn, path string) {
		response := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nWelcome to Sensei!"
		conn.Write([]byte(response))
	})

	server.Router.Handle("GET", "/hello", func(conn net.Conn, path string) {
		response := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nHello, World!"
		conn.Write([]byte(response))
	})

	server.Start()

	//Start the server
	server.Start()

}

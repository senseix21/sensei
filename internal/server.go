package internal

import (
	"bufio"
	"fmt"
	"net"
	"sensei/internal/router"
	"strings"
)

type server struct {
	Addr   string
	Router *router.Router
}

// NewServer creates a new server
func NewServer(addr string) *server {
	return &server{
		Addr:   addr,
		Router: router.NewRouter(),
	}

}

// Start starts the server
func (s *server) Start() error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return fmt.Errorf("failed to start the server: %w", err)
	}

	defer listener.Close()
	fmt.Println("server started on", s.Addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept connection:", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

// handle connection reads the request and sends a response
func (s *server) handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	request, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("failed to read request:", err)
		return
	}

	fmt.Println("received request:", request)

	//parse http request
	parts := strings.Fields(request)
	if len(parts) < 2 {
		fmt.Fprint(conn, "HTTP 1.1 400 Bad Request\r\n\r\n")
		return
	}

	method := parts[0]
	path := parts[1]

	//fmt.Printf("Received request: %s %s\n", method, path)

	//Simple response
	//response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nHello from Sensei! You requested %s %s", method, path)
	//conn.Write([]byte(response))

	//Route the request
	s.Router.Serve(conn, method, path)
}

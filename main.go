package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("new incoming connetion from client:", ws.RemoteAddr())
	s.conns[ws] = true
	s.readloop(ws)
}

func (s *Server) readloop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read err:", err)
			continue
		}
		msg := buf[:n]
		fmt.Println(string(msg))
		ws.Write([]byte("thankyou for the msg!!!"))
		break
	}
}

func main() {
	server := NewServer()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set Content Security Policy header
		w.Header().Set("Content-Security-Policy", "default-src 'self' ws://localhost:3000/ws")
		// Serve your HTML content here
		// For example:
		http.ServeFile(w, r, "index.html")
	})
	http.Handle("/ws", websocket.Handler(server.handleWS))
	fmt.Println("server starting at port :3000")
	http.ListenAndServe(":3000", nil)
}

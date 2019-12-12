package main

import (
	"log"
	"net"
	"net/http"

	"github.com/priyankshah217/producer-api/services"
)

func main() {
	mux := services.GetHttpHandler()
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	log.Printf("API starting: port %d (%s)", 8080, ln.Addr())
	log.Printf("API terminating: %v", http.Serve(ln, mux))
}

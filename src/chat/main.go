package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()

	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	defer listen.Close()
	log.Printf("started server on :8888\n")

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err)
			continue
		}
		go s.newClient(conn)
	}
}

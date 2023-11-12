package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"
	"sync"
)

const (
	HOSTNAME = "localhost"
	PORT     = "5000"
	TYPE     = "tcp"
)

var (
	wg = &sync.WaitGroup{}
)

func main() {

	log.Printf("starting listening at %s:%s", HOSTNAME, PORT)
	listener, err := net.Listen(TYPE, HOSTNAME+":"+PORT)

	if err != nil {
		panic("server cannot start...")
	}
	log.Print("server started")

	for {

		connection, err := listener.Accept()
		defer connection.Close()

		if err != nil {
			log.Print(err)
		}

		go handleConnection(connection)

	}
}

func handleConnection(connection net.Conn) {
	log.Printf("connected to %s", connection.RemoteAddr())
	bufferedReader := bufio.NewReader(connection)

	defer connection.Close()

	for {

		str, err := bufferedReader.ReadString('\n')

		if err != nil {

			if errors.Is(err, io.EOF) {
				log.Printf("client %s disconnected", connection.RemoteAddr())
				return
			}

			log.Print(err)
			return
		}

		log.Println(str)

	}
}

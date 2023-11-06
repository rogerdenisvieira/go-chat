package main

import (
	"bufio"
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

	log.Printf("starting the listening at %s:%s", HOSTNAME, PORT)
	listener, err := net.Listen(TYPE, HOSTNAME+":"+PORT)

	if err != nil {
		panic("server cannot start...")
	}

	connection, err := listener.Accept()
	defer connection.Close()

	if err != nil {
		panic("cannot accept new connections")
	}

	log.Print("server started")

	wg.Add(1)
	go receive(connection)

	wg.Wait()
	log.Print("server stopped")
}

func receive(connection net.Conn) {
	bufferedReader := bufio.NewReader(connection)
	defer wg.Done()

	for {
		str, err := bufferedReader.ReadString('\n')
		if err != nil {
			connection.Close()
			panic(err)
		}

		log.Println(str)
	}
}

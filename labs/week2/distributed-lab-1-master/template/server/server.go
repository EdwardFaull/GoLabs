package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strconv"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	// TODO: all
	// Deal with an error event.
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	for {
		conn, _ := ln.Accept()
		conns <- conn
	}
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.
}

func handleClient(client *net.Conn, clientid int, msgs chan Message) {
	// TODO: all
	// So long as this connection is alive:
	// Read in new messages as delimited by '\n's
	// Tidy up each message and add it to the messages channel,
	// recording which client it came from.
	isOpen := true
	fmt.Println("Started handling client", clientid)
	reader := bufio.NewReader(*client)
	for isOpen {
		text, err := reader.ReadString('\n')
		if err != nil {
			(*client).Close()
			isOpen = false
			fmt.Println("Closed connection with client", clientid)
		}
		fmt.Println("Received", text, "from client", clientid)
		fmt.Fprintln(*client, "OK")
		msg := Message{
			sender:  clientid,
			message: text,
		}
		msgs <- msg
	}
}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	ln, _ := net.Listen("tcp", *portPtr)

	//TODO Create a Listener for TCP connections on the port given above.

	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)

	//Start accepting connections
	go acceptConns(ln, conns)
	for {
		select {
		case conn := <-conns:
			clientsSize := len(clients)
			fmt.Println("Clients:", clientsSize)
			clients[clientsSize] = conn
			go handleClient(&conn, clientsSize, msgs)
			//TODO Deal with a new connection
			// - assign a client ID
			// - add the client to the clients channel
			// - start to asynchronously handle messages from this client
		case msg := <-msgs:
			for i, client := range clients {
				if i != msg.sender {
					sender := strconv.Itoa(msg.sender)
					fmt.Fprintf(client, "Client "+sender+": "+msg.message)
				}
			}
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender
		}
	}
}

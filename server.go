package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func main() {
	go server()
	go client()
	var input string
	fmt.Scanln(&input)
}

func server() {
	//listen on port

	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		//accept connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//handle the connection
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	//receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("received", msg)
	}
	c.Close()
}

func client() {
	//connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	//send message
	msg := "Hello World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

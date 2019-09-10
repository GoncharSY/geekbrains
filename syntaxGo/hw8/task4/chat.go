package main

import (
	"fmt"
	"log"
	"net"
	"os"

	cln "./client"
	srv "./server"
)

var progName = os.Args[0]
var progArgs = os.Args[1:]
var addr = "localhost:8000"

func main() {
	if len(progArgs) == 0 {
		fmt.Println("Не указан режим (client/server)")
		return
	}

	switch progArgs[0] {
	case "server":
		startServer()
	case "client":
		startClient()
	}
}

func startServer() {
	var listener net.Listener
	var connection net.Conn
	var err error
	var chat = srv.NewChat()

	fmt.Println("Запускаю сервер.")
	listener, err = net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Запускаю чат.")
	chat.Start()

	for {
		connection, err = listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		chat.IncomeChannel <- connection
	}
}

func startClient() {
	fmt.Println("Запускаю клиент чата.")
	cln.Start(addr)
}

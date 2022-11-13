package main

import (
	"context"
	"flag"
	"go-backend-1/lesson-02/task-03/client"
	"go-backend-1/lesson-02/task-03/server"
	"log"
	"os"
	"os/signal"
	"strings"
)

var mode string = "server"
var host string = server.DefaultHost
var port string = server.DefaultPort

func main() {
	flag.StringVar(&mode, "mode", mode, "Execution mode: client or server")
	flag.StringVar(&host, "host", host, "Application host")
	flag.StringVar(&port, "port", port, "Application port")
	flag.Parse()

	switch mode {
	case "client":
		startClient()
	case "server":
		startServer()
	default:
		log.Fatal("Unknown mode: ", mode)
	}
}

// Запустить новый игровой-клиент.
func startClient() {
	var ctx, _ = signal.NotifyContext(context.Background(), os.Interrupt)
	var cli = client.New(host, port)
	var adr = cli.GetAddress()

	log.Println("Client starting:", adr)
	log.Println(strings.Repeat("=", 50))

	if err := cli.Start(); err != nil {
		log.Fatal(err)
	}

	select {
	case <-ctx.Done():
		cli.Stop()
	case <-cli.Clossing:
		log.Println("Connection closed by server")
	}

	log.Println(strings.Repeat("=", 50))
	log.Println("Client stopped")
}

// Запустить новый игровой-сервер.
func startServer() {
	var ctx, _ = signal.NotifyContext(context.Background(), os.Interrupt)
	var srv = server.New(host, port)

	log.Println("Server starting:", srv.GetAddress())
	log.Println(strings.Repeat("=", 50))

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	} else {
		<-ctx.Done()
	}

	srv.Stop()
	log.Println(strings.Repeat("=", 50))
	log.Println("Server stopped")
}

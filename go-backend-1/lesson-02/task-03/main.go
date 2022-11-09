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

// Запустить новый чат-клиент.
func startClient() {
	var ctx, _ = signal.NotifyContext(context.Background(), os.Interrupt)
	var cli = client.New(host, port)
	var adr = cli.GetAddress()

	if err := cli.Start(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("client started:", adr)
		log.Println(strings.Repeat("=", 50))
	}

	select {
	case <-ctx.Done():
		cli.Stop()
	case <-cli.Clossing:
		log.Println("connection closed by server")
	}

	log.Println(strings.Repeat("=", 50))
	log.Println("client stopped")
}

// Запустить новый чат-сервер.
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

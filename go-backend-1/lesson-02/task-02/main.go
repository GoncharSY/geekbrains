package main

import (
	"context"
	"flag"
	"go-backend-1/lesson-02/task-02/client"
	"go-backend-1/lesson-02/task-02/server"
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

func startServer() {
	var ctx, _ = signal.NotifyContext(context.Background(), os.Interrupt)
	var srv = server.New(host, port)

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("server started:", srv.GetAddress())
		log.Println(strings.Repeat("=", 50))
	}

	<-ctx.Done()

	if err := srv.Stop(); err != nil {
		log.Fatal(err)
	} else {
		log.Println(strings.Repeat("=", 50))
		log.Println("server stopped")
	}
}

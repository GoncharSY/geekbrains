package main

import (
	"context"
	"flag"
	"fmt"
	"go-backend-1/lesson-02/task-01/client"
	"go-backend-1/lesson-02/task-01/server"
	"log"
	"os"
	"os/signal"
	"strings"
)

const DefaultHost = "localhost"
const DefaultPort = "9000"

func main() {
	var mode string

	flag.StringVar(&mode, "mode", "server", "Execution mode: client or server")
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
	var addr = fmt.Sprintf("%s:%s", DefaultHost, DefaultPort)

	log.Println("client starting with address: ", addr)
	log.Println(strings.Repeat("=", 50))

	if err := client.Start(addr, ctx); err != nil {
		log.Println(err)
	}

	log.Println(strings.Repeat("=", 50))
	log.Println("client stopped")
}

func startServer() {
	var ctx, _ = signal.NotifyContext(context.Background(), os.Interrupt)
	var srv = server.New(DefaultHost, DefaultPort, ctx)

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("server started with address: ", srv.GetAddress())
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

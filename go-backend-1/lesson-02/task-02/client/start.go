package client

// import (
// 	"context"
// 	"fmt"
// 	"log"
// )

// func Start(host, port string, ctx context.Context) *Structure {
// 	var cli = New(host, port)

// 	if err := cli.Start(); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		defer cli.Stop()
// 	}

// 	select {
// 	case <-ctx.Done():
// 		return nil
// 	case msg := <-cli.Receiving:
// 		log.Println(msg)
// 	case msg := <-cli.Sending:
// 		fmt.Fprintln(cli.Connection, msg)
// 	}

// 	return cli
// }

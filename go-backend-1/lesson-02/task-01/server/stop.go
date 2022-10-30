package server

import (
	"log"
	"net"
)

// Остановить сервер.
func (srv *Structure) Stop() error {
	srv.Cancel()
	srv.ConnectGroup.Wait()
	log.Println("all connections are closed")
	return nil
}

// Прекратить обработку соединения.
func (srv *Structure) StopConnect(con net.Conn) {
	con.Close()
	log.Println("connection closed", con.RemoteAddr())
	srv.ConnectGroup.Done()
}

package server

import (
	"log"
)

// Остановить сервер.
func (srv *Structure) Stop() error {
	go srv.StopConnections()

	srv.Cancel()
	srv.ConnectGroup.Wait()

	log.Println("all connections are closed")
	return nil
}

// Остановить (разорвать) все соединения с клиентами.
func (srv *Structure) StopConnections() {
	for addr := range srv.Connections {
		srv.RemoveConnection(addr)
	}
}

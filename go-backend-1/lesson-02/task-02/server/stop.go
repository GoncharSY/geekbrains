package server

import "log"

// Остановить чат-сервер.
// При этом все соединения с клиентами будут закрыты.
func (srv *Structure) Stop() error {
	srv.Cancel()
	srv.ClientGroup.Wait()

	log.Println("all connections are closed")

	close(srv.Entering)
	close(srv.Leaving)
	close(srv.Messaging)

	return nil
}

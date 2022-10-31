package server

import "log"

func (srv *Structure) Stop() error {
	srv.Cancel()
	srv.ClientGroup.Wait()
	log.Println("all connections are closed")

	close(srv.Entering)
	close(srv.Leaving)
	close(srv.Messaging)

	return nil
}

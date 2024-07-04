package player

import "net"

// Создать нового игрока на сервере и вернуть указатель на него.
func New(conn net.Conn) *Structure {
	var plr = &Structure{
		Name:       conn.RemoteAddr().String(),
		Connection: conn,
		Answering:  make(chan string, 10),
	}

	go plr.startAnswering()

	return plr
}

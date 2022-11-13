package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	"go-backend-1/lesson-02/task-03/server/player"
	"go-backend-1/lesson-02/task-03/server/quest"
)

// Структура для игрового сервера.
type Structure struct {
	Host string
	Port string

	Context context.Context
	Cancel  context.CancelFunc

	Config   net.ListenConfig
	Listener net.Listener

	Players  map[string]*player.Structure
	Question *quest.Structure
	QuestMtx sync.Mutex

	Broadcast chan string
}

// ======================================================================================
// ======================================================================================
// ======================================================================================

// Получить адрес, на котром запускается сервер.
func (srv *Structure) GetAddress() string {
	return fmt.Sprintf("%v:%v", srv.Host, srv.Port)
}

// Отправить сообщение всем игрокам на сервере.
func (srv *Structure) Announce(msg string) {
	for _, plr := range srv.Players {
		go plr.Send(msg)
	}
}

// ======================================================================================
// STOP =================================================================================
// ======================================================================================

// Остановить работу сервера.
func (srv *Structure) Stop() {
	srv.Cancel()
}

// ======================================================================================
// START ================================================================================
// ======================================================================================

// Запустить сервер. При этом автоматически начнется новая игра.
func (srv *Structure) Start() error {
	var ctx = srv.Context
	var adr = srv.GetAddress()
	var err error

	srv.Listener, err = srv.Config.Listen(ctx, "tcp", adr)

	if err == nil {
		srv.startNewQuest()
		go srv.startProcessing()
		go srv.startAccepting()
	}

	return err
}

// Начать обработку.
func (srv *Structure) startProcessing() {
	defer log.Println("processing stopped")

	for {
		select {
		case <-srv.Context.Done():
			return
		case msg := <-srv.Broadcast:
			srv.Announce(msg)
		}
	}
}

// Начать подключение новых игроков.
func (srv *Structure) startAccepting() {
	for {
		con, err := srv.Listener.Accept()

		if err != nil {
			log.Println("acceptance error", err)
			return
		} else {
			go srv.startPlay(con)
		}
	}
}

// Начать взаимодействие с игроком.
func (srv *Structure) startPlay(conn net.Conn) {
	var plr = player.New(conn)
	var err error

	if err = srv.askName(plr); err != nil {
		log.Println("naming error:", err)
		return
	} else if err = srv.addPlayer(plr); err != nil {
		log.Println("joining error:", err)
		return
	} else {
		defer srv.removePlayer(plr.Name)
		defer log.Println(plr.Name, "stopped playing")
	}

	log.Println("Player " + plr.Name + " joined")
	plr.Send(fmt.Sprintf(FormatOfWelcome, plr.Name))
	srv.askAnswer(plr)

	for {
		select {
		case <-srv.Context.Done():
			return
		case msg, opn := <-plr.Answering:
			if !opn {
				return
			} else {
				go srv.acceptAnswer(msg, plr)
			}
		}
	}
}

// Начать новую игру.
func (srv *Structure) startNewQuest() {
	srv.Question.Reset()
	srv.Announce(fmt.Sprintf(FormatOfAnswerRequest, srv.Question))
	log.Println("New game started:", srv.Question)
}

// ======================================================================================
// PLAYERS ==============================================================================
// ======================================================================================

// Проверить наличие игрока под указанным именем.
func (srv *Structure) hasPlayer(name string) bool {
	_, yes := srv.Players[name]
	return yes
}

// Удалить игрока.
func (srv *Structure) removePlayer(name string) {
	if srv.hasPlayer(name) {
		delete(srv.Players, name)
	}
}

// Добавить нового игрока.
func (srv *Structure) addPlayer(p *player.Structure) error {
	if srv.hasPlayer(p.Name) {
		return errors.New("already exists")
	}

	srv.Players[p.Name] = p
	return nil
}

// Запросить у игрока его имя.
func (srv *Structure) askName(p *player.Structure) error {
	for {
		p.Send("Enter your name please:\n")

		select {
		case <-srv.Context.Done():
			return errors.New("game canceled")
		case name, open := <-p.Answering:
			if !open {
				return errors.New("player left")
			} else if srv.hasPlayer(name) {
				p.Send(fmt.Sprintf(FormatOfPlayerExists, name))
				continue
			} else {
				p.Name = name
				return nil
			}
		}
	}
}

// ======================================================================================
// QUESTION AND ANSWER ==================================================================
// ======================================================================================

// Проверить корректность ответа на игровую задачу.
func (srv *Structure) isCorrect(ans string) bool {
	if val, err := strconv.Atoi(ans); err != nil {
		return false
	} else {
		return srv.Question.IsSolution(val)
	}
}

// Принять ответ игрока для проверки его на корректность.
func (srv *Structure) acceptAnswer(ans string, plr *player.Structure) {
	srv.QuestMtx.Lock()
	defer srv.QuestMtx.Unlock()

	if srv.isCorrect(ans) {
		log.Print(plr.Name + " won!\n")
		srv.Announce(plr.Name + " won!\n")
		time.Sleep(time.Second)
		srv.Announce("New game is starting...\n")
		time.Sleep(time.Second)
		srv.startNewQuest()
	} else {
		plr.Send("Your answer is wrong!\n")
		srv.askAnswer(plr)
	}
}

// Отправить игроку условия актуальной игровой задачи.
func (srv *Structure) askAnswer(plr *player.Structure) {
	plr.Send(fmt.Sprintf(FormatOfAnswerRequest, srv.Question))
}

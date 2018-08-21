package robot

import (
	"github.com/mingz2013/chinese-chess-table-go/chess"
	"github.com/mingz2013/chinese-chess-table-go/msg"
	"log"
	"time"
)

type Robot struct {
	Id   int
	Name string

	TableId int

	MsgIn  <-chan msg.Msg
	MsgOut chan<- msg.Msg

	chess.ChessBoard
}

func (r *Robot) Init() {
	r.ChessBoard.Init()
}

func NewRobot(id int, name string, msgIn <-chan msg.Msg, msgOut chan<- msg.Msg) Robot {
	r := Robot{Id: id, Name: name, MsgIn: msgIn, MsgOut: msgOut}
	r.Init()
	return r
}

func (r Robot) Run() {

	r.doSit()

	for {

		select {
		case m, ok := <-r.MsgIn:
			{
				if !ok {
					continue
				}

				r.onMsg(m)
			}
		case <-time.After(1 * time.Second):
			continue
		}

	}
}

func (r *Robot) onMsg(m msg.Msg) {
	switch m.GetCmd() {
	//case "table":
	//	r.onTableMsg(m)
	//case "play":
	//	r.onPlayMsg(m)
	default:
		log.Println("unknown msg", m)

	}
}

func (r Robot) doSit() {

}

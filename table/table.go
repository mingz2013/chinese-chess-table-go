package table

import (
	"github.com/mingz2013/chinese-chess-table-go/msg"
	"github.com/mingz2013/chinese-chess-table-go/player"
	"log"
	"time"
)

type Table struct {
	Id int

	MsgIn  <-chan msg.Msg
	MsgOut chan<- msg.Msg

	PlayersManager
}

func NewTable(id int, msgIn <-chan msg.Msg, msgOut chan<- msg.Msg) Table {
	t := Table{Id: id, MsgIn: msgIn, MsgOut: msgOut}
	t.Init()
	return t
}

func (t *Table) Init() {
	for i := 0; i < 4; i++ {
		t.Players[i] = player.NewPlayer(i)
	}
	//t.Play = NewPlay(t)

}

func (t Table) Run() {

	for {
		select {
		case m, ok := <-t.MsgIn:
			{
				if !ok {
					continue
				}

				t.onMsg(m)
			}
		case <-time.After(1 * time.Second):
			continue

		}

	}
}

func (t *Table) onMsg(m msg.Msg) {
	log.Println("Table.onMsg", m, &t, &t.Players)
	switch m.GetCmd() {
	case "table":
		t.onTableMsg(m)
	//case "play":
	//	t.Play.OnMsg(m)
	default:
		log.Println("unknown cmd", m)
	}
}

func (t *Table) onTableMsg(m msg.Msg) {
	params := m.GetParams()
	action := params["action"].(string)
	switch action {
	case "sit":
		//t.onTableSit(m)
	}

}

package table

import (
	"github.com/mingz2013/chinese-chess-table-go/chess"
	"github.com/mingz2013/chinese-chess-table-go/msg"
	"log"
	"time"
)

type Table struct {
	Id int

	MsgIn  <-chan msg.Msg
	MsgOut chan<- msg.Msg

	TablePlayers

	*chess.ChessBoard
}

func NewTable(id int, msgIn <-chan msg.Msg, msgOut chan<- msg.Msg) Table {
	t := Table{Id: id, MsgIn: msgIn, MsgOut: msgOut}
	t.Init()
	return t
}

func (t *Table) Init() {
	t.TablePlayers.Init()
	t.ChessBoard = chess.NewChessBoard()

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
	case "play":
		t.onPlayMsg(m)
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

func (t *Table) onPlayMsg(m msg.Msg) {
	params := m.GetParams()
	action := params["action"].(string)
	switch action {
	case "move":
		action_ := &chess.Action{}
		action_.ParseFromMsg(params)
		// TODO 验证是否有权限doaction，比如，player是否对应相应颜色。当前是否该这个人走了
		t.ChessBoard.DoAction(action_)
	default:
		log.Println("Table.onPlayMsg", "unknown action", action)

	}
}

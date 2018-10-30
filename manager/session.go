package manager

import (
	"github.com/mingz2013/chinese-chess-table-go/table"
	"github.com/mingz2013/lib-go/msg"
)

type UserSession struct {
	UserId  int
	TableId int
}

func NewUserSession(userId int) UserSession {
	return UserSession{UserId: userId}
}

type TableSession struct {
	Table   table.Table
	MsgIn   chan msg.Msg
	MsgOut  chan msg.Msg
	TableId int
}

func NewTableSession(table table.Table, msgIn chan msg.Msg, msgOut chan msg.Msg, tableId int) TableSession {
	return TableSession{Table: table, MsgIn: msgIn, MsgOut: msgOut, TableId: tableId}
}

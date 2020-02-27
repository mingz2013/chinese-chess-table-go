package main

import (
	"github.com/mingz2013/chinese-chess-table-go/board"
	"log"
)

func getActionList() (actions []board.Action) {

	actions = append(actions,
		board.NewAction2(0, 0, 0, 1),
		board.NewAction2(0, 1, 0, 2),
		board.NewAction2(0, 2, 0, 3),
		board.NewAction2(0, 3, 0, 4),
		board.NewAction2(0, 4, 0, 5),
		board.NewAction2(0, 5, 0, 6),
		board.NewAction2(0, 6, 0, 7),
	)

	return
}

func main() {

	actionList := getActionList()

	b := board.NewChessBoard()

	for _, action := range actionList {
		ok := b.DoAction(action)
		log.Println("for action ok", action, ok)
		if !ok {
			break
		}
	}
}

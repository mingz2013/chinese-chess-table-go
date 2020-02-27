package main

import (
	"github.com/mingz2013/chinese-chess-table-go/board"
	"log"
)

func getActionList() (actions []board.Action) {

	actions = append(actions,
		board.NewAction2(0, 0, 0, 1), // 車
		board.NewAction2(0, 1, 0, 2), // 車

		board.NewAction2(1, 0, 2, 2), // 马

		board.NewAction2(1, 9, 0, 7),

		board.NewAction2(2, 0, 4, 2),

		board.NewAction2(3, 0, 4, 1),
		board.NewAction2(4, 0, 3, 0),
	)

	return
}

func main() {

	actionList := getActionList()

	b := board.NewChessBoard()

	for _, action := range actionList {
		log.Println("for action begin", action)
		ok := b.DoAction(action)
		if !ok {
			log.Println("for action error", action.Src.X(), action.Src.Y(), action.Dst.X(), action.Dst.Y())
			break
		}
	}
}

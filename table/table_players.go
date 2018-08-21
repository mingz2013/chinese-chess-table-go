package table

import "github.com/mingz2013/chinese-chess-table-go/player"

type TablePlayers struct {
	Players [2]player.Player
}

func (t *TablePlayers) Init() {
	for i := 0; i < 4; i++ {
		t.Players[i] = player.NewPlayer(i)
	}
}

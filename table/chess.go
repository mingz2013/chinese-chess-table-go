package table

const (
	CHESS_JU = iota
	CHESS_MA
	CHESS_XIANG
	CHESS_SHI
	CHESS_SHUAI
	CHESS_PAO
	CHESS_BING

	COLOR_RED   = 0
	COLOR_BLACK = 1
)

type Chess struct {
	color int
	cType int
}

func NewChess(color, cType int) *Chess {
	return &Chess{color: color, cType: cType}
}

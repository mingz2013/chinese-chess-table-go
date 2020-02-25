package chinese_chess_table_go

const (
	CHESS_NONE = iota
	CHESS_JU
	CHESS_MA
	CHESS_XIANG
	CHESS_SHI
	CHESS_SHUAI
	CHESS_PAO
	CHESS_BING

	COLOR_NONE  = 0
	COLOR_RED   = 1
	COLOR_BLACK = 2
)

type chess struct {
	color int
	cType int
}

func NewChess(color, cType int) *chess {
	return &chess{color: color, cType: cType}
}

func (c *chess) IsNone() bool {
	return c.color == COLOR_NONE || c.cType == CHESS_NONE
}

func (c *chess) Color() int {
	return c.color
}

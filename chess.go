package chinese_chess_table_go

const (
	CHESS_NONE  = 0
	CHESS_JU    = 1
	CHESS_MA    = 2
	CHESS_XIANG = 3
	CHESS_SHI   = 4
	CHESS_SHUAI = 5
	CHESS_PAO   = 6
	CHESS_BING  = 7

	COLOR_NONE  = 0
	COLOR_RED   = 1
	COLOR_BLACK = 2
)

// | ctype 4bit | color 4bit |
type Chess uint8

func (c Chess) cType() uint8 {
	return uint8(c) >> 4
}

func (c Chess) color() uint8 {
	return uint8(c) << 4 >> 4
}

func NewChess(color, cType uint8) (c Chess) {
	i := color + cType<<4
	return Chess(i)
}

func (c Chess) isNone() bool {
	return c.color() == COLOR_NONE || c.cType() == CHESS_NONE
	//return uint8(c) == 0
}

func otherColor(color uint8) uint8 {
	if color == COLOR_BLACK {
		return COLOR_RED
	} else if color == COLOR_RED {
		return COLOR_BLACK
	} else {
		panic("error color")
	}
}

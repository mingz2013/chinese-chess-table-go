package table

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

type Chess struct {
	color int
	cType int
}

func NewChess(color, cType int) *Chess {
	return &Chess{color: color, cType: cType}
}

func (c *Chess) IsNone() bool {
	return c.color == COLOR_NONE || c.cType == CHESS_NONE
}

func (c *Chess) Color() int {
	return c.color
}

func (c *Chess) DoAction() bool {

}

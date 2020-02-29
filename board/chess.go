package board

import (
	"fmt"
	"log"
)

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
	return (uint8(c) << 4) >> 4
}

func (c Chess) isNone() bool {
	return c.color() == COLOR_NONE || c.cType() == CHESS_NONE
	//return uint8(c) == 0
}

func (c Chess) cTypeString() string {
	switch c.cType() {
	case CHESS_NONE:
		return "空"
	case CHESS_JU:
		return "車"
	case CHESS_MA:
		return "马"
	case CHESS_XIANG:
		return "象"
	case CHESS_SHI:
		return "士"
	case CHESS_SHUAI:
		return "帅"
	case CHESS_PAO:
		return "炮"
	case CHESS_BING:
		return "兵"
	default:
		log.Fatal(c.cType())
		return ""
	}

}

func (c Chess) colorString() string {
	switch c.color() {
	case COLOR_NONE:
		return "空"
	case COLOR_RED:
		return "红"
	case COLOR_BLACK:
		return "黑"
	default:
		log.Fatal(c.color())
		return ""
	}
}

func (c Chess) String() string {
	//return fmt.Sprintf("Chess(%v, %v)", c.colorString(), c.cTypeString())

	if c.color() == COLOR_RED {
		return fmt.Sprintf("（%v）", c.cTypeString())
	}

	return fmt.Sprintf("【%v】", c.cTypeString())
}

func NewChess(color, cType uint8) (c Chess) {
	i := (cType << 4) + color
	return Chess(i)
}

func otherColor(color uint8) uint8 {
	log.Println("otherColor <<", color)
	if color == COLOR_BLACK {
		return COLOR_RED
	} else if color == COLOR_RED {
		return COLOR_BLACK
	} else {
		panic("error color")
	}
}

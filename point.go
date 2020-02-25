package chinese_chess_table_go

// | x 4bit | y 4bit |
type Point uint8

func (p Point) X() uint8 {
	return uint8(p) >> 4
}

func (p Point) Y() uint8 {
	return uint8(p) << 4 >> 4
}

func NewPoint(x, y uint8) (p Point) {
	i := x + y<<4
	return Point(i)
}

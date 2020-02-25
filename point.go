package chinese_chess_table_go

type point struct {
	X int
	Y int
}

func NewPoint(x, y int) *point {
	return &point{x, y}
}

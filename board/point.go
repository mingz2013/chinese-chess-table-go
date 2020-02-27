package board

// 9 * 10
type Point int8

func (p Point) X() int8 {
	return int8(p) % 9
}

func (p Point) Y() int8 {
	return int8(p) / 9
}

func (a Point) getInfo() {

}

func NewPoint(x, y int8) (p Point) {
	return Point(x + y*9)
}

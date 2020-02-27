package board

type Action struct {
	Src Point
	Dst Point
}

func (a Action) getInfo() {

}

func NewAction(src, dst Point) Action {
	return Action{src, dst}
}

func NewAction2(sx, sy, dx, dy int8) Action {
	return Action{NewPoint(sx, sy), NewPoint(dx, dy)}
}

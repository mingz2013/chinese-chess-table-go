package board

import "fmt"

type Action struct {
	Src Point
	Dst Point
}

func (a Action) getInfo() {

}
func (a Action) String() string {
	return fmt.Sprintf("Action(%v, %v)", a.Src, a.Dst)
}

func NewAction(src, dst Point) Action {
	return Action{src, dst}
}

func NewAction2(sx, sy, dx, dy int8) Action {
	return Action{NewPoint(sx, sy), NewPoint(dx, dy)}
}

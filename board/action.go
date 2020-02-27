package board

import "fmt"

type Action struct {
	Src Point
	Dst Point
}

func (a Action) getInfo() {

}
func (c Action) String() string {
	return fmt.Sprintf("Action(%v, %v)", c.Src, c.Dst)
}

func NewAction(src, dst Point) Action {
	return Action{src, dst}
}

func NewAction2(sx, sy, dx, dy int8) Action {
	return Action{NewPoint(sx, sy), NewPoint(dx, dy)}
}

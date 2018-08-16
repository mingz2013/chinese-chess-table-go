package table

type Action struct {
	Src Point
	Dst Point
}

func NewAction(src, dst Point) *Action {
	return &Action{src, dst}
}

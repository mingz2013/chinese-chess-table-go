package chinese_chess_table_go

type action struct {
	Src *point
	Dst *point
}

func NewAction(src, dst *point) *action {
	return &action{src, dst}
}

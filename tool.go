package chinese_chess_table_go

func Abs(a int8) (ret int8) {
	ret = (a ^ a>>7) - a>>7
	return
}

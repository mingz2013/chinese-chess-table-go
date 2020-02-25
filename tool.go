package chinese_chess_table_go

func Abs(a int) (ret int) {
	ret = (a ^ a>>31) - a>>31
	return
}

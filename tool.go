package chinese_chess_table_go

func Abs(a uint8) (ret uint8) {
	ret = (a ^ a>>31) - a>>31
	return
}

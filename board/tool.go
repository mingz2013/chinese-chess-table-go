package board

func Abs(i int8) int8 {
	return (i ^ i>>7) - i>>7
}

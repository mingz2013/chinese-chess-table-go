package table

// 我方为红方视角，以红方左下侧为原点

type ChessBoard [9][10]*Chess

func (c *ChessBoard) Init() {
	c[0][0] = NewChess(CHESS_JU, COLOR_RED)
	c[1][0] = NewChess(CHESS_MA, COLOR_RED)
	c[2][0] = NewChess(CHESS_XIANG, COLOR_RED)
	c[3][0] = NewChess(CHESS_SHI, COLOR_RED)
	c[4][0] = NewChess(CHESS_SHUAI, COLOR_RED)
	c[5][0] = NewChess(CHESS_SHI, COLOR_RED)
	c[6][0] = NewChess(CHESS_XIANG, COLOR_RED)
	c[7][0] = NewChess(CHESS_MA, COLOR_RED)
	c[8][0] = NewChess(CHESS_JU, COLOR_RED)

	c[2][2] = NewChess(CHESS_PAO, COLOR_RED)
	c[8][2] = NewChess(CHESS_PAO, COLOR_RED)

	c[0][3] = NewChess(CHESS_BING, COLOR_RED)
	c[2][3] = NewChess(CHESS_BING, COLOR_RED)
	c[4][3] = NewChess(CHESS_BING, COLOR_RED)
	c[6][3] = NewChess(CHESS_BING, COLOR_RED)
	c[8][3] = NewChess(CHESS_BING, COLOR_RED)

}

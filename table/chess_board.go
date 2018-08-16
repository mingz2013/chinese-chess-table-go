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

	c[0][6] = NewChess(CHESS_BING, COLOR_BLACK)
	c[2][6] = NewChess(CHESS_BING, COLOR_BLACK)
	c[4][6] = NewChess(CHESS_BING, COLOR_BLACK)
	c[6][6] = NewChess(CHESS_BING, COLOR_BLACK)
	c[8][6] = NewChess(CHESS_BING, COLOR_BLACK)

	c[2][7] = NewChess(CHESS_PAO, COLOR_BLACK)
	c[8][7] = NewChess(CHESS_PAO, COLOR_BLACK)

	c[0][9] = NewChess(CHESS_JU, COLOR_BLACK)
	c[1][9] = NewChess(CHESS_MA, COLOR_BLACK)
	c[2][9] = NewChess(CHESS_XIANG, COLOR_BLACK)
	c[3][9] = NewChess(CHESS_SHI, COLOR_BLACK)
	c[4][9] = NewChess(CHESS_SHUAI, COLOR_BLACK)
	c[5][9] = NewChess(CHESS_SHI, COLOR_BLACK)
	c[6][9] = NewChess(CHESS_XIANG, COLOR_BLACK)
	c[7][9] = NewChess(CHESS_MA, COLOR_BLACK)
	c[8][9] = NewChess(CHESS_JU, COLOR_BLACK)

}

func (c *ChessBoard) GetInfo() {

}

func (c *ChessBoard) GetChessBYPoint(point Point) *Chess {
	return c.GetChess(point.X, point.Y)
}

func (c *ChessBoard) GetChess(x, y int) *Chess {
	return c[x][y]
}

func (c *ChessBoard) SetChess(x int, y int, chess *Chess) {
	c[x][y] = chess
}

func (c *ChessBoard) ClearChess(x, y int) {
	c[x][y] = nil
}

func (c *ChessBoard) DoAction(action Action, src, dst *Chess) {
	c.SetChess(action.Dst.X, action.Dst.Y, src)
	c.ClearChess(action.Src.X, action.Src.Y)
}

func (c *ChessBoard) CheckAction(action Action) (ok bool) {
	srcChess := c.GetChessBYPoint(action.Src)
	dstChess := c.GetChessBYPoint(action.Dst)

	ok = c.CheckActionSameColor(action, srcChess, dstChess)
	if !ok {
		return false
	}

	switch srcChess.cType {
	case CHESS_NONE:
		return false
	case CHESS_JU:
		ok = c.CheckJuAction(action, srcChess, dstChess)
	case CHESS_MA:
		ok = c.CheckMaAction(action, srcChess, dstChess)
	case CHESS_XIANG:
		ok = c.CheckXiangAction(action, srcChess, dstChess)
	case CHESS_SHI:
		ok = c.CheckShiAction(action, srcChess, dstChess)
	case CHESS_SHUAI:
		ok = c.CheckShuaiAction(action, srcChess, dstChess)
	case CHESS_PAO:
		ok = c.CheckPaoAction(action, srcChess, dstChess)
	case CHESS_BING:
		ok = c.CheckBingAction(action, srcChess, dstChess)
	default:
		return false
	}
	return
}

func (c *ChessBoard) CheckActionSameColor(action Action, srcChess, dstChess *Chess) bool {
	// 相同颜色的不能互相吃
	if srcChess.Color() == dstChess.Color() {
		return false
	}
	return true
}

func (c *ChessBoard) CheckJuAction(action Action, srcChess, dstChess *Chess) bool {
	// 检查路径，是否满足路线规则

	// 在一条直线上，

	// 检查路径上是否有碍事的

	if action.Src.Y != action.Dst.Y && action.Src.X != action.Dst.X {
		return false
	}

	if action.Src.Y == action.Dst.Y {
		tmp := 1
		if action.Src.X-action.Dst.X > 0 {
			tmp = 1
		} else {
			tmp = -1
		}

		for i := action.Src.X + tmp; i < action.Dst.X; i += tmp {
			if !c.GetChess(i, action.Src.Y).IsNone() {
				return false
			}

		}
	} else if action.Src.X == action.Dst.X {
		tmp := 1
		if action.Src.Y-action.Dst.Y > 0 {
			tmp = 1
		} else {
			tmp = -1
		}

		for i := action.Src.Y + tmp; i < action.Dst.Y; i += tmp {
			if !c.GetChess(action.Src.X, i).IsNone() {
				return false
			}

		}
	}

	return true

}

func (c *ChessBoard) CheckMaAction(action Action, srcChess, dstChess *Chess) bool {
	return false
}

func (c *ChessBoard) CheckXiangAction(action Action, srcChess, dstChess *Chess) bool {
	return false
}

func (c *ChessBoard) CheckShiAction(action Action, srcChess, dstChess *Chess) bool {
	return false
}

func (c *ChessBoard) CheckShuaiAction(action Action, srcChess, dstChess *Chess) bool {
	return false
}

func (c *ChessBoard) CheckPaoAction(action Action, srcChess, dstChess *Chess) bool {
	return false
}

func (c *ChessBoard) CheckBingAction(action Action, srcChess, dstChess *Chess) bool {
	return false
}

package chinese_chess_table_go

import "log"

// 我方为红方视角，以红方左下侧为原点

type chessBoard [9][10]*Chess

func NewChessBoard() *chessBoard {
	c := &chessBoard{}
	c.init()
	return c
}

func (c *chessBoard) init() {
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

func (c *chessBoard) getInfo() {
}

func (c *chessBoard) getChessBYPoint(point *Point) *Chess {
	return c.getChess(point.X, point.Y)
}

func (c *chessBoard) getChess(x, y int) *Chess {
	return c[x][y]
}

func (c *chessBoard) setChess(x int, y int, chess *Chess) {
	c[x][y] = chess
}

func (c *chessBoard) clearChess(x, y int) {
	c[x][y] = nil
}

func (c *chessBoard) DoAction(action *Action) (ok bool) {

	ok = c.checkAction(action)
	if !ok {
		return
	}

	srcChess := c.getChessBYPoint(action.Src)
	//dstChess := c.GetChessBYPoint(action.Dst)

	c.setChess(action.Dst.X, action.Dst.Y, srcChess)
	c.clearChess(action.Src.X, action.Src.Y)

	return
}

func (c *chessBoard) checkAction(action *Action) (ok bool) {
	// 检查规则，action是否可以正确执行

	srcChess := c.getChessBYPoint(action.Src)
	dstChess := c.getChessBYPoint(action.Dst)

	ok = c.checkActionSameColor(action, srcChess, dstChess)
	if !ok {
		return false
	}

	switch srcChess.cType {
	case CHESS_NONE:
		return false
	case CHESS_JU:
		ok = c.checkJuAction(action, srcChess, dstChess)
	case CHESS_MA:
		ok = c.checkMaAction(action, srcChess, dstChess)
	case CHESS_XIANG:
		ok = c.checkXiangAction(action, srcChess, dstChess)
	case CHESS_SHI:
		ok = c.checkShiAction(action, srcChess, dstChess)
	case CHESS_SHUAI:
		ok = c.checkShuaiAction(action, srcChess, dstChess)
	case CHESS_PAO:
		ok = c.checkPaoAction(action, srcChess, dstChess)
	case CHESS_BING:
		ok = c.checkBingAction(action, srcChess, dstChess)
	default:
		return false
	}
	return
}

func (c *chessBoard) checkActionSameColor(action *Action, srcChess, dstChess *Chess) bool {
	// 相同颜色的不能互相吃
	if srcChess.Color() == dstChess.Color() {
		return false
	}
	return true
}

func (c *chessBoard) checkJuAction(action *Action, srcChess, dstChess *Chess) bool {
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
			if !c.getChess(i, action.Src.Y).IsNone() {
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
			if !c.getChess(action.Src.X, i).IsNone() {
				return false
			}

		}
	}

	return true

}

func (c *chessBoard) checkMaAction(action *Action, srcChess, dstChess *Chess) bool {

	// 马走日，

	// 别马脚

	if Abs(action.Src.X-action.Dst.X) == 1 {
		if Abs(action.Src.Y-action.Dst.Y) != 2 {
			return false
		}
	} else if Abs(action.Src.X-action.Dst.Y) == 2 {
		if Abs(action.Src.Y-action.Dst.Y) != 1 {
			return false
		}
	} else {
		return false
	}

	var maJiaoPoint *Point
	if Abs(action.Src.X-action.Dst.X) == 1 {
		maJiaoPoint = NewPoint(action.Src.X, (action.Src.Y-action.Dst.Y)/2+action.Src.Y)

	} else {
		maJiaoPoint = NewPoint((action.Src.X-action.Dst.X)/2+action.Src.X, action.Src.Y)
	}

	if !c.getChessBYPoint(maJiaoPoint).IsNone() {
		return false
	}

	return false
}

func (c *chessBoard) checkXiangAction(action *Action, srcChess, dstChess *Chess) bool {
	// 象走田
	// 别象眼

	if Abs(action.Src.X-action.Dst.X) != 2 || Abs(action.Src.Y-action.Dst.Y) != 2 {
		return false
	}

	xiangYanPoint := NewPoint((action.Src.X-action.Dst.X)/2+action.Src.X, (action.Src.Y-action.Dst.Y)/2+action.Src.Y)
	if !c.getChessBYPoint(xiangYanPoint).IsNone() {
		return false
	}

	return false
}

func (c *chessBoard) checkShiAction(action *Action, srcChess, dstChess *Chess) bool {
	//士，
	// 只能斜着走
	// 只能在九宫格

	if Abs(action.Src.X-action.Dst.X) != 1 || Abs(action.Src.Y-action.Dst.Y) != 1 {
		return false
	}

	// 锁定中间三个坐标
	if action.Dst.X != 3 && action.Dst.X != 4 && action.Dst.X != 5 {
		return false
	}

	if srcChess.Color() == COLOR_RED {
		if action.Dst.Y != 0 && action.Dst.Y != 1 && action.Dst.Y != 2 {
			return false
		}
	} else if srcChess.Color() == COLOR_BLACK {
		if action.Dst.Y != 7 && action.Dst.Y != 8 && action.Dst.Y != 9 {
			return false
		}
	} else {
		log.Println("err...")
		return false
	}

	return true
}

func (c *chessBoard) checkShuaiAction(action *Action, srcChess, dstChess *Chess) bool {

	// 帅，只能在九宫格

	// 在一条直线上
	if action.Src.Y != action.Dst.Y && action.Src.X != action.Dst.X {
		return false
	}
	// 只能移动一位
	if Abs(action.Src.X-action.Dst.X)+Abs(action.Src.Y-action.Dst.Y) != 1 {
		return false
	}

	// 锁定中间三个坐标
	if action.Dst.X != 3 && action.Dst.X != 4 && action.Dst.X != 5 {
		return false
	}

	if srcChess.Color() == COLOR_RED {
		if action.Dst.Y != 0 && action.Dst.Y != 1 && action.Dst.Y != 2 {
			return false
		}
	} else if srcChess.Color() == COLOR_BLACK {
		if action.Dst.Y != 7 && action.Dst.Y != 8 && action.Dst.Y != 9 {
			return false
		}
	} else {
		log.Println("err...")
		return false
	}

	return true
}

func (c *chessBoard) checkPaoAction(action *Action, srcChess, dstChess *Chess) bool {

	// 炮
	// 横竖移动多个位置，
	// 可以跳过一个打对方一个

	// 一条直线上
	if action.Src.Y != action.Dst.Y && action.Src.X != action.Dst.X {
		return false
	}

	// 检查中间有几个位置有子

	centerCount := 0

	if action.Src.Y == action.Dst.Y {
		tmp := 1
		if action.Src.X-action.Dst.X > 0 {
			tmp = 1
		} else {
			tmp = -1
		}

		for i := action.Src.X + tmp; i < action.Dst.X; i += tmp {
			if !c.getChess(i, action.Src.Y).IsNone() {
				centerCount += 1
				if centerCount > 1 {
					return false
				}

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
			if !c.getChess(action.Src.X, i).IsNone() {
				centerCount += 1
				if centerCount > 1 {
					return false
				}
			}

		}
	}

	if centerCount == 1 {
		if dstChess == nil || dstChess.IsNone() {
			return false
		}
	} else if centerCount == 0 {
		if dstChess != nil && !dstChess.IsNone() {
			return false
		}
	}

	return true
}

func (c *chessBoard) checkBingAction(action *Action, srcChess, dstChess *Chess) bool {
	// 兵，过河之前，过河之后，
	// 过河之前只能向前，过河之后可以向前和平移

	// 在一条直线上
	if action.Src.Y != action.Dst.Y && action.Src.X != action.Dst.X {
		return false
	}
	// 只能移动一位
	if Abs(action.Src.X-action.Dst.X)+Abs(action.Src.Y-action.Dst.Y) != 1 {
		return false
	}

	if srcChess.Color() == COLOR_RED {
		if action.Src.Y <= 4 {
			// 未过河的红色小卒

			if action.Src.X != action.Dst.X || action.Src.Y+1 != action.Dst.Y {
				return false
			}

		} else {
			// 过河的卒子

			if action.Src.Y > action.Dst.Y {
				return false
			}

		}
	} else if srcChess.Color() == COLOR_BLACK {
		if action.Src.Y >= 5 {
			if action.Src.X != action.Dst.X || action.Src.Y-1 != action.Dst.Y {
				return false
			}
		} else {
			if action.Src.Y < action.Dst.Y {
				return false
			}
		}
	} else {
		log.Println("err..")
		return false
	}

	return true
}

func (c *chessBoard) checkJiangJun(color int) (ok bool) {
	// 检查将军

	var shuaiPoint *Point
	if color == COLOR_RED {
		shuaiPoint = c.getRedShuaiPoint()
	} else {
		shuaiPoint = c.getBlackShuaiPoint()
	}

	// 直接遍历对方所有的棋，看哪个棋可以doaction到帅的位置

	for x := 0; x < 9; x++ {
		for y := 0; y < 10; y++ {
			chess := c.getChess(x, y)
			if chess != nil && chess.Color() != color {
				point := NewPoint(x, y)
				ok = c.checkAction(NewAction(point, shuaiPoint))
				if ok {
					return true
				}
			}
		}
	}

	return false
}

func (c *chessBoard) getRedShuaiPoint() (shuaiPoint *Point) {

	for i := 3; i <= 5; i++ {
		for j := 0; j <= 2; j++ {
			chess := c.getChess(i, j)
			if chess != nil && chess.cType == CHESS_SHUAI {
				return NewPoint(i, j)
			}
		}
	}

	return nil
}

func (c *chessBoard) getBlackShuaiPoint() (shuaiPoint *Point) {

	for i := 3; i <= 5; i++ {
		for j := 7; j <= 9; j++ {
			chess := c.getChess(i, j)
			if chess != nil && chess.cType == CHESS_SHUAI {
				return NewPoint(i, j)
			}
		}
	}

	return nil
}

func (c *chessBoard) checkWin() {
	// 检查输赢
}

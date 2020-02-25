package chinese_chess_table_go

import "log"

// 我方为红方视角，以红方左下侧为原点

//type chessBoard [9][10]Chess
type chessBoard [9 * 10]Chess

func NewChessBoard() *chessBoard {
	c := &chessBoard{}
	c.init()
	return c
}

func (c *chessBoard) init() {
	c[NewPoint(0, 0)] = NewChess(CHESS_JU, COLOR_RED)
	c[NewPoint(1, 0)] = NewChess(CHESS_MA, COLOR_RED)
	c[NewPoint(2, 0)] = NewChess(CHESS_XIANG, COLOR_RED)
	c[NewPoint(3, 0)] = NewChess(CHESS_SHI, COLOR_RED)
	c[NewPoint(4, 0)] = NewChess(CHESS_SHUAI, COLOR_RED)
	c[NewPoint(5, 0)] = NewChess(CHESS_SHI, COLOR_RED)
	c[NewPoint(6, 0)] = NewChess(CHESS_XIANG, COLOR_RED)
	c[NewPoint(7, 0)] = NewChess(CHESS_MA, COLOR_RED)
	c[NewPoint(8, 0)] = NewChess(CHESS_JU, COLOR_RED)

	c[NewPoint(2, 2)] = NewChess(CHESS_PAO, COLOR_RED)
	c[NewPoint(8, 2)] = NewChess(CHESS_PAO, COLOR_RED)

	c[NewPoint(0, 3)] = NewChess(CHESS_BING, COLOR_RED)
	c[NewPoint(2, 3)] = NewChess(CHESS_BING, COLOR_RED)
	c[NewPoint(4, 3)] = NewChess(CHESS_BING, COLOR_RED)
	c[NewPoint(6, 3)] = NewChess(CHESS_BING, COLOR_RED)
	c[NewPoint(8, 3)] = NewChess(CHESS_BING, COLOR_RED)

	c[NewPoint(0, 6)] = NewChess(CHESS_BING, COLOR_BLACK)
	c[NewPoint(2, 6)] = NewChess(CHESS_BING, COLOR_BLACK)
	c[NewPoint(4, 6)] = NewChess(CHESS_BING, COLOR_BLACK)
	c[NewPoint(6, 6)] = NewChess(CHESS_BING, COLOR_BLACK)
	c[NewPoint(8, 6)] = NewChess(CHESS_BING, COLOR_BLACK)

	c[NewPoint(2, 7)] = NewChess(CHESS_PAO, COLOR_BLACK)
	c[NewPoint(8, 7)] = NewChess(CHESS_PAO, COLOR_BLACK)

	c[NewPoint(0, 9)] = NewChess(CHESS_JU, COLOR_BLACK)
	c[NewPoint(1, 9)] = NewChess(CHESS_MA, COLOR_BLACK)
	c[NewPoint(2, 9)] = NewChess(CHESS_XIANG, COLOR_BLACK)
	c[NewPoint(3, 9)] = NewChess(CHESS_SHI, COLOR_BLACK)
	c[NewPoint(4, 9)] = NewChess(CHESS_SHUAI, COLOR_BLACK)
	c[NewPoint(5, 9)] = NewChess(CHESS_SHI, COLOR_BLACK)
	c[NewPoint(6, 9)] = NewChess(CHESS_XIANG, COLOR_BLACK)
	c[NewPoint(7, 9)] = NewChess(CHESS_MA, COLOR_BLACK)
	c[NewPoint(8, 9)] = NewChess(CHESS_JU, COLOR_BLACK)

}

func (c *chessBoard) getInfo() {
}

func (c *chessBoard) getChessByPoint(point Point) Chess {
	return c[point]
}

func (c *chessBoard) getChess(x, y int8) Chess {
	return c.getChessByPoint(NewPoint(x, y))
}

func (c *chessBoard) setChess(point Point, chess Chess) {
	c[point] = chess
}

func (c *chessBoard) clearChess(point Point) {
	c[point] = 0
}

func (c *chessBoard) move(action Action) (dstChess Chess) {
	srcChess := c.getChessByPoint(action.Src)
	dstChess = c.getChessByPoint(action.Dst)

	c.setChess(action.Dst, srcChess)
	c.clearChess(action.Src)

	return dstChess
}

func (c *chessBoard) rollbackMove(action Action, dstChess Chess) {
	srcChess := c.getChessByPoint(action.Src)
	dstChess = c.getChessByPoint(action.Dst)

	c.setChess(action.Src, srcChess)
	c.clearChess(action.Dst)
	if dstChess != 0 {
		c.setChess(action.Dst, dstChess)
	}
}

func (c *chessBoard) testMove(action Action) bool {
	// 测试走动
	//dstChess := c.getChessBYPoint(action.Dst)
	srcChess := c.getChessByPoint(action.Src)

	dstChess := c.move(action)
	// 走动后是否被将军
	isJiang := c.checkJiangJun(srcChess.color())

	c.rollbackMove(action, dstChess)

	return isJiang

}

func (c *chessBoard) DoAction(action Action) (ok bool) {

	// 检查棋子走动是否符合规则
	ok = c.checkAction(action)
	if !ok {
		return
	}

	// 检查棋子走动后，是否会被将军
	ok = c.testMove(action)
	if !ok {
		return
	}

	// 走动棋子
	c.move(action)

	// 检查棋子走动后是否将对方

	// 检查棋子走动后是否赢
	c.checkWin(c.getChessByPoint(action.Dst).color())
	return
}

func (c *chessBoard) checkAction(action Action) (ok bool) {
	// 检查规则，action是否可以正确执行

	srcChess := c.getChessByPoint(action.Src)
	dstChess := c.getChessByPoint(action.Dst)

	ok = c.checkActionSameColor(action, srcChess, dstChess)
	if !ok {
		return false
	}

	switch srcChess.cType() {
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

func (c *chessBoard) checkActionSameColor(action Action, srcChess, dstChess Chess) bool {
	// 相同颜色的不能互相吃
	if srcChess.color() == dstChess.color() {
		return false
	}
	return true
}

func (c *chessBoard) checkJuAction(action Action, srcChess, dstChess Chess) bool {
	// 检查路径，是否满足路线规则

	// 在一条直线上，

	// 检查路径上是否有碍事的

	if action.Src.Y() != action.Dst.Y() && action.Src.X() != action.Dst.X() {
		return false
	}

	if action.Src.Y() == action.Dst.Y() {
		var tmp int8
		if action.Src.X()-action.Dst.X() > 0 {
			tmp = 1
		} else {
			tmp = -1
		}

		for i := action.Src.X() + tmp; i < action.Dst.X(); i += tmp {
			if !c.getChess(i, action.Src.Y()).isNone() {
				return false
			}

		}
	} else if action.Src.X() == action.Dst.X() {
		var tmp int8
		if action.Src.Y()-action.Dst.Y() > 0 {
			tmp = 1
		} else {
			tmp = -1
		}

		for i := action.Src.Y() + tmp; i < action.Dst.Y(); i += tmp {
			if !c.getChess(action.Src.X(), i).isNone() {
				return false
			}

		}
	}

	return true

}

func (c *chessBoard) checkMaAction(action Action, srcChess, dstChess Chess) bool {

	// 马走日，

	// 别马脚

	if Abs(action.Src.X()-action.Dst.X()) == 1 {
		if Abs(action.Src.Y()-action.Dst.Y()) != 2 {
			return false
		}
	} else if Abs(action.Src.X()-action.Dst.Y()) == 2 {
		if Abs(action.Src.Y()-action.Dst.Y()) != 1 {
			return false
		}
	} else {
		return false
	}

	var maJiaoPoint Point
	if Abs(action.Src.X()-action.Dst.X()) == 1 {
		maJiaoPoint = NewPoint(action.Src.X(), action.Src.Y()-action.Dst.Y()/2+action.Src.Y())

	} else {
		maJiaoPoint = NewPoint((action.Src.X()-action.Dst.X())/2+action.Src.X(), action.Src.Y())
	}

	if !c.getChessByPoint(maJiaoPoint).isNone() {
		return false
	}

	return false
}

func (c *chessBoard) checkXiangAction(action Action, srcChess, dstChess Chess) bool {
	// 象走田
	// 别象眼

	if Abs(action.Src.X()-action.Dst.X()) != 2 || Abs(action.Src.Y()-action.Dst.Y()) != 2 {
		return false
	}

	xiangYanPoint := NewPoint((action.Src.X()-action.Dst.X())/2+action.Src.X(), (action.Src.Y()-action.Dst.Y())/2+action.Src.Y())
	if !c.getChessByPoint(xiangYanPoint).isNone() {
		return false
	}

	return false
}

func (c *chessBoard) checkShiAction(action Action, srcChess, dstChess Chess) bool {
	//士，
	// 只能斜着走
	// 只能在九宫格

	if Abs(action.Src.X()-action.Dst.X()) != 1 || Abs(action.Src.Y()-action.Dst.Y()) != 1 {
		return false
	}

	// 锁定中间三个坐标
	if action.Dst.X() != 3 && action.Dst.X() != 4 && action.Dst.X() != 5 {
		return false
	}

	if srcChess.color() == COLOR_RED {
		if action.Dst.Y() != 0 && action.Dst.Y() != 1 && action.Dst.Y() != 2 {
			return false
		}
	} else if srcChess.color() == COLOR_BLACK {
		if action.Dst.Y() != 7 && action.Dst.Y() != 8 && action.Dst.Y() != 9 {
			return false
		}
	} else {
		log.Println("err...")
		return false
	}

	return true
}

func (c *chessBoard) checkShuaiAction(action Action, srcChess, dstChess Chess) bool {

	// 帅，只能在九宫格

	// 在一条直线上
	if action.Src.Y() != action.Dst.Y() && action.Src.X() != action.Dst.X() {
		return false
	}
	// 只能移动一位
	if Abs(action.Src.X()-action.Dst.X())+Abs(action.Src.Y()-action.Dst.Y()) != 1 {
		return false
	}

	// 锁定中间三个坐标
	if action.Dst.X() != 3 && action.Dst.X() != 4 && action.Dst.X() != 5 {
		return false
	}

	if srcChess.color() == COLOR_RED {
		if action.Dst.Y() != 0 && action.Dst.Y() != 1 && action.Dst.Y() != 2 {
			return false
		}
	} else if srcChess.color() == COLOR_BLACK {
		if action.Dst.Y() != 7 && action.Dst.Y() != 8 && action.Dst.Y() != 9 {
			return false
		}
	} else {
		log.Println("err...")
		return false
	}

	return true
}

func (c *chessBoard) checkPaoAction(action Action, srcChess, dstChess Chess) bool {

	// 炮
	// 横竖移动多个位置，
	// 可以跳过一个打对方一个

	// 一条直线上
	if action.Src.Y() != action.Dst.Y() && action.Src.X() != action.Dst.X() {
		return false
	}

	// 检查中间有几个位置有子

	centerCount := 0

	if action.Src.Y() == action.Dst.Y() {
		var tmp int8
		if action.Src.X()-action.Dst.X() > 0 {
			tmp = 1
		} else {
			tmp = -1
		}

		for i := action.Src.X() + tmp; i < action.Dst.X(); i += tmp {
			if !c.getChess(i, action.Src.Y()).isNone() {
				centerCount += 1
				if centerCount > 1 {
					return false
				}

			}

		}
	} else if action.Src.X() == action.Dst.X() {
		var tmp int8
		if action.Src.Y()-action.Dst.Y() > 0 {
			tmp = 1
		} else {
			tmp = -1
		}

		for i := action.Src.Y() + tmp; i < action.Dst.Y(); i += tmp {
			if !c.getChess(action.Src.X(), i).isNone() {
				centerCount += 1
				if centerCount > 1 {
					return false
				}
			}

		}
	}

	if centerCount == 1 {
		if dstChess == 0 || dstChess.isNone() {
			return false
		}
	} else if centerCount == 0 {
		if dstChess != 0 && !dstChess.isNone() {
			return false
		}
	}

	return true
}

func (c *chessBoard) checkBingAction(action Action, srcChess, dstChess Chess) bool {
	// 兵，过河之前，过河之后，
	// 过河之前只能向前，过河之后可以向前和平移

	// 在一条直线上
	if action.Src.Y() != action.Dst.Y() && action.Src.X() != action.Dst.X() {
		return false
	}
	// 只能移动一位
	if Abs(action.Src.X()-action.Dst.X())+Abs(action.Src.Y()-action.Dst.Y()) != 1 {
		return false
	}

	if srcChess.color() == COLOR_RED {
		if action.Src.Y() <= 4 {
			// 未过河的红色小卒

			if action.Src.X() != action.Dst.X() || action.Src.Y()+1 != action.Dst.Y() {
				return false
			}

		} else {
			// 过河的卒子

			if action.Src.Y() > action.Dst.Y() {
				return false
			}

		}
	} else if srcChess.color() == COLOR_BLACK {
		if action.Src.Y() >= 5 {
			if action.Src.X() != action.Dst.X() || action.Src.Y()-1 != action.Dst.Y() {
				return false
			}
		} else {
			if action.Src.Y() < action.Dst.Y() {
				return false
			}
		}
	} else {
		log.Println("err..")
		return false
	}

	return true
}

func (c *chessBoard) checkJiangJun(color uint8) (ok bool) {
	// 检查将军, color, 被将军的颜色

	var shuaiPoint Point
	if color == COLOR_RED {
		shuaiPoint = c.getRedShuaiPoint()
	} else {
		shuaiPoint = c.getBlackShuaiPoint()
	}

	// 直接遍历对方所有的棋，看哪个棋可以doaction到帅的位置
	var x, y int8
	for x = 0; x < 9; x++ {
		for y = 0; y < 10; y++ {
			chess := c.getChess(x, y)
			if chess != 0 && chess.color() != color {
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

func (c *chessBoard) getRedShuaiPoint() (shuaiPoint Point) {
	var x, y int8
	for x = 3; x <= 5; x++ {
		for y = 0; y <= 2; y++ {
			chess := c.getChess(x, y)
			if chess != 0 && chess.cType() == CHESS_SHUAI {
				return NewPoint(x, y)
			}
		}
	}

	return 0
}

func (c *chessBoard) getBlackShuaiPoint() (shuaiPoint Point) {
	var x, y int8
	for x = 3; x <= 5; x++ {
		for y = 7; y <= 9; y++ {
			chess := c.getChess(x, y)
			if chess != 0 && chess.cType() == CHESS_SHUAI {
				return NewPoint(x, y)
			}
		}
	}

	return 0
}

func (c *chessBoard) getAllCanActionChess(p Point) (actions []Action) {
	// 找出chess所有可以移动的action
	var x, y int8
	for x = 0; x < 9; x++ {
		for y = 0; y < 10; y++ {
			newPoint := NewPoint(x, y)
			newAction := NewAction(p, newPoint)
			ok := c.checkAction(newAction)
			if ok {
				actions = append(actions, newAction)
			}
		}
	}

	return
}

func (c *chessBoard) getAllCanAction(color uint8) (actions []Action) {
	// 获取color方的所有可以移动的action
	var x, y int8
	for x = 0; x < 9; x++ {
		for y = 0; y < 10; y++ {
			chess := c.getChess(x, y)
			if chess != 0 && chess.color() == color {
				point := NewPoint(x, y)
				actionsTmp := c.getAllCanActionChess(point)
				actions = append(actions, actionsTmp...) // slice合并
			}
		}
	}

	return
}

func (c *chessBoard) checkWin(color uint8) bool {
	// 检查color方是否赢
	// 检查输赢，我方走了一步，
	// 最笨的方法，遍历对方可以走动的步骤，然后检查是否仍然将军
	if !c.checkJiangJun(otherColor(color)) {
		return false
	}

	actions := c.getAllCanAction(otherColor(color))

	for _, action := range actions {
		if !c.testMove(action) {
			// 如果走动后没有被将军，则没有赢
			return false
		}
	}

	return true

}

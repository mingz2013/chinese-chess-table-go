package chinese_chess_table_go

func getActionList() (actions []Action) {
	return
}

func Main() {

	actionList := getActionList()

	b := NewChessBoard()

	for _, action := range actionList {
		ok := b.DoAction(action)
		if !ok {
			break
		}
	}
}

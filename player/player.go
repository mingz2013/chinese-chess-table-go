package player

type Player struct {
	Id   int
	Name string

	ChessColor int // 红还是 黑

}

func (p *Player) Init() {
	//p.Cards = cards.NewCards()
}

func NewPlayer(chessColor int) Player {

	p := Player{ChessColor: chessColor}
	p.Init()
	return p
}

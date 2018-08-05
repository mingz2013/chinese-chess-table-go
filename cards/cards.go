package cards

type Cards struct {
	nowTile int
}

func (c *Cards) Init() {
}

func NewCards() Cards {
	return Cards{}
}

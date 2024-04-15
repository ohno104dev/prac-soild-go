package bad

type Bird interface {
	Fly() string
}

type Pigeon struct{}

func (p *Pigeon) Fly() string {
	return "I can fly"
}

type Penguin struct{}

// contradicts the expected behavior of a bird
func (p *Penguin) Fly() string {
	return "I can't fly"
}

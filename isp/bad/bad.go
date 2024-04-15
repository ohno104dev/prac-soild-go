package bad

// Not all animals can fly or swim
type Animal interface {
	Walk() string
	Fly() string
	Swim() string
}

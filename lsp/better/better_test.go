package better

import "testing"

func TestBetter(t *testing.T) {
	LetItWalk(&Pigeon{})
	LetItWalk(&Penguin{})

	LetItFly(&Pigeon{})

	// LetItFly(&Penguin{}) //This doesn't compile, which is good
}

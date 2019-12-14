package main

import (
	"fmt"
	"gitlab.com/read-music-learner/musical-notation/domain"
	"gitlab.com/read-music-learner/musical-notation/domain/pitch"
)

func main() {
	fmt.Print(domain.Note{
		Pitch:      pitch.C,
		Octave:     1,
		Duration:   0,
		Alteration: pitch.Flat,
	}.Print())
}

package main

import (
	"fmt"
	"github.com/happy-developer-fr/musical_notation/musical_notation"
	"github.com/happy-developer-fr/musical_notation/musical_notation/pitch"
)

func main() {
	fmt.Println(musical_notation.Note{
		Pitch:      pitch.C,
		Octave:     1,
		Duration:   0,
		Alteration: pitch.Flat,
	}.Print())
}

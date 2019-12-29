package main

import (
	"fmt"
	"gitlab.com/read-music-learner/musical-notation/musical_notation"
	"gitlab.com/read-music-learner/musical-notation/musical_notation/pitch"
)

func main() {
	fmt.Println(musical_notation.Note{
		Pitch:      pitch.C,
		Octave:     1,
		Duration:   0,
		Alteration: pitch.Flat,
	}.Print())
}

package main

import (
	"fmt"
	"github.com/happy-developer-fr/musical-notation/pkg/note"
	"github.com/happy-developer-fr/musical-notation/pkg/pitch"
)

func main() {
	fmt.Println(note.Note{
		Pitch:      pitch.C,
		Octave:     1,
		Duration:   0,
		Alteration: pitch.Flat,
	}.Print())
}

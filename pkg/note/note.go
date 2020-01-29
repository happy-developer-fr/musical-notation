package note

import (
	"fmt"
	"github.com/happy-developer-fr/musical-notation/pkg/pitch"
	"github.com/happy-developer-fr/musical-notation/pkg/rhythm"
)

//Note Define a note structure
type Note struct {
	Pitch      pitch.Pitch
	Octave     pitch.Octave
	Alteration pitch.Alteration
	Duration   rhythm.Duration
}

//Print Simply print to string a note
func (n Note) Print() string {
	return fmt.Sprint(n.Pitch, n.Octave.Print(), n.Alteration, n.Duration)
}

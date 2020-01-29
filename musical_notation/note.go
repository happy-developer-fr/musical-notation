package musical_notation

import (
	"fmt"

	"github.com/happy-developer-fr/musical_notation/musical_notation/pitch"
	"github.com/happy-developer-fr/musical_notation/musical_notation/rhythm"
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

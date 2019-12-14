package domain

import (
	"fmt"
	"gitlab.com/read-music-learner/musical-notation/domain/pitch"
	"gitlab.com/read-music-learner/musical-notation/domain/rhythm"
)

type Note struct {
	Pitch      pitch.Pitch
	Octave     pitch.Octave
	Alteration pitch.Alteration
	Duration   rhythm.Duration
}

func (n Note) Print() string {
	return fmt.Sprint(n.Pitch, n.Octave.Print(), n.Alteration, n.Duration)
}

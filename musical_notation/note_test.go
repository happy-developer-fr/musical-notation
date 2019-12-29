package musical_notation

import (
	"gitlab.com/read-music-learner/musical-notation/musical_notation/pitch"
	"testing"
)

func TestNoteFormatWithDieseAndPositiveOctave(t *testing.T) {
	note := Note{
		Pitch:      pitch.C,
		Octave:     2,
		Alteration: pitch.Sharp,
		Duration:   0.25,
	}

	expecting := "C''#0.25"
	if note.Print() != expecting {
		t.Error("Format is incorrect, expecting", expecting)
	}
}

func TestNoteFormatWithBemolAndNegativeOctave(t *testing.T) {
	note := Note{
		Pitch:      pitch.C,
		Octave:     -2,
		Alteration: pitch.Flat,
		Duration:   0.25,
	}
	expecting := "C,,b0.25"
	if note.Print() != expecting {
		t.Error("Format is incorrect, expecting", expecting)
	}
}

package song

import (
	"gitlab.com/read-music-learner/musical-notation/musical_notation"
	"gitlab.com/read-music-learner/musical-notation/musical_notation/pitch"
	"testing"
)

func TestSong(t *testing.T) {
	note1 := musical_notation.Note{
		Pitch:      pitch.C,
		Octave:     2,
		Alteration: pitch.Sharp,
		Duration:   0.25,
	}
	note2 := musical_notation.Note{
		Pitch:      pitch.C,
		Octave:     2,
		Alteration: pitch.Sharp,
		Duration:   0.25,
	}

	songNotes := []musical_notation.Note{note1, note2}
	if len(songNotes) != 2 {
		t.Error("Song should have two songNotes")
	}

	song := Song{notes: songNotes}

	if song.notes[0] != note1 {
		t.Error("First note should be first")
	}

	if song.notes[1] != note2 {
		t.Error("First note should be first")
	}
}

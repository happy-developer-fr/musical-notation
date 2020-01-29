package song

import (
	"github.com/happy-developer-fr/musical-notation/pkg/note"
	"github.com/happy-developer-fr/musical-notation/pkg/pitch"
	"testing"
)

var note1 = note.Note{
	Pitch:      pitch.C,
	Octave:     2,
	Alteration: pitch.Sharp,
	Duration:   0.25,
}
var note2 = note.Note{
	Pitch:    pitch.D,
	Octave:   0,
	Duration: 1,
}
var songNotes = []note.Note{note1, note2}
var twoNoteSong = Song{Notes: songNotes}

func TestSong(t *testing.T) {

	if len(twoNoteSong.Notes) != 2 {
		t.Error("Song should have two songNotes")
	}

	if twoNoteSong.Notes[0] != note1 {
		t.Error("First note should be first")
	}

	if twoNoteSong.Notes[1] != note2 {
		t.Error("First note should be first")
	}

	songPrinted := twoNoteSong.Print()
	songExpected := "C''#0.25, D1"
	if songPrinted != songExpected {
		t.Error("Wrong twoNoteSong, printed : ", songPrinted, "expected:", songExpected)
	}
}

func TestSerializeSong(t *testing.T) {
	songJson, err := twoNoteSong.PrintJson()
	if err != nil {
		t.Error("an error occured when marshalling a song")
	}
	expectedSongJson := `{"Notes":[{"Pitch":"C","Octave":2,"Alteration":"#","Duration":0.25},{"Pitch":"D","Octave":0,"Alteration":"","Duration":1}]}`
	if songJson == "{}" || songJson == "" {
		t.Error("Song should not be empty")
	}

	if songJson != expectedSongJson {
		t.Error("Song:", songJson, "not well formatted")
	}
}

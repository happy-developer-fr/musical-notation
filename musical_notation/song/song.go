package song

import (
	"encoding/json"
	"github.com/happy-developer-fr/musical_notation/musical_notation"
)

//Song Represent a list of Notes
type Song struct {
	Notes []musical_notation.Note
}

//Concat all notes
func (s Song) Print() string {
	var songPrinted = ""
	for i, note := range s.Notes {
		if i != 0 {
			songPrinted += ", "
		}
		songPrinted += note.Print()
	}
	return songPrinted
}

func (s Song) PrintJson() (string, error) {
	b, err := json.Marshal(s)

	return string(b), err
}

package song

import (
	"encoding/json"
	"github.com/happy-developer-fr/musical-notation/pkg/note"
	uuid "github.com/satori/go.uuid"
)

//Song Represent a list of Notes
type Song struct {
	Id    uuid.UUID
	Notes []note.Note
}

//Concat all notes
func (s Song) Print() string {
	var songPrinted = ""
	for i, aNote := range s.Notes {
		if i != 0 {
			songPrinted += ", "
		}
		songPrinted += aNote.Print()
	}
	return songPrinted
}

func (s Song) PrintJson() (string, error) {
	b, err := json.Marshal(s)

	return string(b), err
}

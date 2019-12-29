package song

import "gitlab.com/read-music-learner/musical-notation/musical_notation"

//Song Represent a list of notes
type Song struct {
	notes []musical_notation.Note
}

func (s Song) print() string {
	var songPrinted = ""
	for i, note := range s.notes {
		if i != 0 {
			songPrinted += ", "
		}
		songPrinted += note.Print()
	}
	return songPrinted
}

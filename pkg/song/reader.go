package song

import (
	"github.com/happy-developer-fr/musical-notation/pkg/note"
)

type Reader interface {
	Next() (note.Note, uint32)
}

type songReader struct {
	Song     Song
	Position uint32
}

func (self *songReader) Next() (note.Note, uint32) {
	n := self.Song.Notes[self.Position]
	indice := self.Position
	self.increment()
	return n, indice
}

func (self *songReader) increment() {
	self.Position++
}

func NewReader(s Song) Reader {
	return &songReader{s, 0}
}

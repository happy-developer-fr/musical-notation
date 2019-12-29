package song

import "gitlab.com/read-music-learner/musical-notation/musical_notation"

type Reader interface {
	Next() (musical_notation.Note, uint32)
}

type songReader struct {
	Song     Song
	Position uint32
}

func (self *songReader) Next() (musical_notation.Note, uint32) {
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

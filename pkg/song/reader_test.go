package song

import (
	"testing"
)

var reader = NewReader(twoNoteSong)

func Test_songReader_Next(t *testing.T) {

	if n, i := reader.Next(); n != note1 || i != 0 {
		t.Error("First note is not the expected one:", n, i)
	}

	if n, i := reader.Next(); n != note2 || i != 1 {
		t.Error("Second note is not the expected one", n, i)
	}
}

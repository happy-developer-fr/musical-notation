package musical_notation

import (
	"fmt"
	"github.com/happy-developer-fr/musical_notation/musical_notation/pitch"
	"testing"
)

func TestRandomizeNote(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%v ", RandomNote([]pitch.Pitch{pitch.C, pitch.G}).Print())
	}
}

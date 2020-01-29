package note

import (
	"fmt"
	"github.com/happy-developer-fr/musical-notation/pkg/pitch"
	"testing"
)

func TestRandomizeNote(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%v ", RandomNote([]pitch.Pitch{pitch.C, pitch.G}).Print())
	}
}

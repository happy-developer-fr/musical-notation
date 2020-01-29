package musical_notation

import (
	"github.com/happy-developer-fr/musical_notation/musical_notation/pitch"
	"github.com/happy-developer-fr/musical_notation/musical_notation/rhythm"
	"math/rand"
	"sync"
	"time"
)

var doOnce sync.Once

func randomPitch(pitches []pitch.Pitch) pitch.Pitch {
	randomValue := rand.Intn(len(pitches))
	return pitches[randomValue]
}

func randomAlteration() pitch.Alteration {
	randomValue := rand.Intn(3)
	switch randomValue {
	case 0:
		return pitch.Sharp
	case 1:
		return pitch.Flat
	default:
		return ""
	}
}

func randomDuration() rhythm.Duration {
	randomValue := rand.Intn(4)
	switch randomValue {
	case 0:
		return 1
	case 2:
		return 0.5
	case 3:
		return 0.25
	default:
		return 0.125

	}
}

func RandomNote(pitches []pitch.Pitch) Note {
	doOnce.Do(func() {
		rand.Seed(time.Now().UTC().UnixNano())
	})
	randomPitch := randomPitch(pitches)
	randomAlteration := randomAlteration()
	duration := randomDuration()
	return Note{Pitch: randomPitch, Alteration: randomAlteration, Duration: duration}
}

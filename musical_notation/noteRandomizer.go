package musical_notation

import (
	"gitlab.com/read-music-learner/musical-notation/musical_notation/pitch"
	"math/rand"
	"sync"
	"time"
)

var doOnce sync.Once

func randomPitch(pitches []pitch.Pitch) pitch.Pitch {
	randomValue := rand.Intn(len(pitches))
	return pitches[randomValue]
}

func RandomNote(pitches []pitch.Pitch) Note {
	doOnce.Do(func() {
		rand.Seed(time.Now().UTC().UnixNano())
	})
	randomPitch := randomPitch(pitches)
	return Note{Pitch: randomPitch}
}

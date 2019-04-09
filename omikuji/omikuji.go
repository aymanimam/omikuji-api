package omikuji

import (
	"fmt"
	"github.com/aymanimam/omikuji-api/errors"
	"math/rand"
	"sync"
	"time"
)

// Randomizer Interface to get random Omikuji from the predefined set of Omikujis
type Randomizer interface {
	GetRandom(min, max int) Omikuji
	GetMax() int
	GetDaikichiMin() int
	GetNoDaikichiMin() int
}

// Omikuji text definition
type Omikuji struct {
	Text string `json:"omikuji"`
}

// AllOmikujis all predefined omikujis
type AllOmikujis []Omikuji

// GetRandom get a random omikuji
func (omikujis *AllOmikujis) GetRandom(min, max int) Omikuji {
	if min < 0 || max > omikujis.GetMax() || min >= max {
		msg := fmt.Sprintf("Invalid arguments: min=%d, max=%d", min, max)
		errors.ThrowOmikujiException(msg, errors.OmikujiErrorCode)
	}

	randIndex := min + rand.Intn(max-min)
	return (*omikujis)[randIndex]
}

// GetMax get the max number of omikujis
func (omikujis *AllOmikujis) GetMax() int {
	return len(*omikujis)
}

// GetDaikichiMin get the min index of all omikujis including the "Daikichi"
func (omikujis *AllOmikujis) GetDaikichiMin() int {
	return 0
}

// GetNoDaikichiMin get the min index of all omikujis excluding the "Daikichi"
func (omikujis *AllOmikujis) GetNoDaikichiMin() int {
	return 1
}

// allOmikujis Singleton all omikujis
var allOmikujis AllOmikujis

// once used for creating a singleton AllOmikujis instance
var once sync.Once

// GetOmikujiRandomizer Get all omikujis singleton instance
func GetOmikujiRandomizer() Randomizer {
	// Using once here for thread safety
	// http://marcio.io/2015/07/singleton-pattern-in-go/
	once.Do(func() {
		// Initialize this var only once
		allOmikujis = []Omikuji{
			{"大吉"},
			{"中吉"},
			{"小吉"},
			{"吉"},
			{"半吉"},
			{"末吉"},
			{"末小吉"},
			{"凶"},
			{"小凶"},
			{"半凶"},
			{"末凶"},
			{"大凶"},
		}
		// Initialize random generator only once
		rand.Seed(time.Now().UTC().UnixNano())
	})
	return &allOmikujis
}

package omikuji

import (
	"testing"
)

var omikujis = []Omikuji{
	{"大吉"},
	{"吉"},
	{"中吉"},
	{"小吉"},
	{"半吉"},
	{"末吉"},
	{"末小吉"},
	{"凶"},
	{"小凶"},
	{"半凶"},
	{"末凶"},
	{"大凶"},
}

func TestGetAllOmikujis(t *testing.T) {
	randomizer := GetOmikujiRandomizer()
	omikujiCount := randomizer.GetMax()
	if omikujiCount != len(omikujis) {
		t.Error(`omikujiCount = `, omikujiCount, `, omikujiCount != 12`)
	}
}

func TestAllOmikujisIsSingleton(t *testing.T) {
	f := func(ch chan Randomizer) {
		randomizer := GetOmikujiRandomizer()
		ch <- randomizer
	}

	randomizerChan1 := make(chan Randomizer)
	randomizerChan2 := make(chan Randomizer)

	go f(randomizerChan1)
	go f(randomizerChan2)

	randomizer1 := <-randomizerChan1
	randomizer2 := <-randomizerChan2

	// Comparing pointers
	if randomizer1 != randomizer2 {
		t.Error(`Two different allOmikujis objects`, randomizer1, randomizer2)
	}
}

func TestGetRandom(t *testing.T) {
	randomizer := GetOmikujiRandomizer()
	omikuji := randomizer.GetRandom(1, 4)
	if !Contains(omikujis[1:4], omikuji) {
		t.Error(`This omikuji [`, omikuji, `] is not expected`)
	}
}

func TestGetRandomInvalidArgs(t *testing.T) {
	AssertPanic(t, "GetOmikujiRandomizer should have panicked!", func() {
		// This function should cause a panic
		randomizer := GetOmikujiRandomizer()
		randomizer.GetRandom(-1, 4)
	})
}

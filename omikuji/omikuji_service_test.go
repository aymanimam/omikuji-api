package omikuji

import (
	"fmt"
	"testing"
	"time"
)

// Mock Randomizer
type MockRandomizer struct{}

func (omikujis *MockRandomizer) GetRandom(min, max int) Omikuji {
	if min == 0 {
		return Omikuji{"大吉"}
	}
	return Omikuji{"吉"}
}

func (omikujis *MockRandomizer) GetMax() int {
	return 2
}

func (omikujis *MockRandomizer) GetDaikichiMin() int {
	return 0
}

func (omikujis *MockRandomizer) GetNoDaikichiMin() int {
	return 1
}

// ---------------

func TestGetPeriodChecker(t *testing.T) {
	currentTime := time.Now()
	currentMonth := currentTime.Month()
	currentDay := currentTime.Day()

	fromDate := PeriodicDate{Month: currentMonth, Day: currentDay}
	toDate := PeriodicDate{Month: currentMonth, Day: currentDay}

	pc := GetPeriodChecker(fromDate, toDate)
	if pc == nil {
		t.Error(`Expected not nil PeriodChecker!`)
	}
}

func TestGetPeriodCheckerInvalidMonthOrder(t *testing.T) {
	AssertPanic(t, "GetPeriodChecker should have panicked!", func() {
		fromDate := PeriodicDate{Month: time.March, Day: 1}
		toDate := PeriodicDate{Month: time.January, Day: 1}
		GetPeriodChecker(fromDate, toDate)
	})
}

func TestGetPeriodCheckerInvalidDayOrder(t *testing.T) {
	AssertPanic(t, "GetPeriodChecker should have panicked!", func() {
		fromDate := PeriodicDate{Month: time.March, Day: 5}
		toDate := PeriodicDate{Month: time.March, Day: 1}
		GetPeriodChecker(fromDate, toDate)
	})
}

func TestPeriod_WithinThePeriod(t *testing.T) {
	fromDate := PeriodicDate{Month: time.January, Day: 1}
	toDate := PeriodicDate{Month: time.March, Day: 1}
	pc := GetPeriodChecker(fromDate, toDate)

	layout := "2006-01-02"
	str := fmt.Sprintf("%d-02-20", time.Now().Year())
	targetTime, _ := time.Parse(layout, str)
	if !pc.WithinThePeriod(targetTime) {
		t.Error(`Expected to be inside the period!`)
	}

	str = fmt.Sprintf("%d-08-20", time.Now().Year())
	targetTime, _ = time.Parse(layout, str)
	if pc.WithinThePeriod(targetTime) {
		t.Error(`Expected to be outside the period!`)
	}

	str = fmt.Sprintf("%d-01-01", time.Now().Year())
	targetTime, _ = time.Parse(layout, str)
	if !pc.WithinThePeriod(targetTime) {
		t.Error(`Expected to be inside the period!`)
	}

	str = fmt.Sprintf("%d-03-01", time.Now().Year())
	targetTime, _ = time.Parse(layout, str)
	if !pc.WithinThePeriod(targetTime) {
		t.Error(`Expected to be inside the period!`)
	}
}

func TestGetOmikujiDispatcher(t *testing.T) {
	fromDate := PeriodicDate{Month: time.January, Day: 1}
	toDate := PeriodicDate{Month: time.March, Day: 1}
	pc := GetPeriodChecker(fromDate, toDate)

	dispatcher := GetOmikujiDispatcher(pc, &MockRandomizer{})
	if dispatcher == nil {
		t.Error(`Dispatcher is expected not to be nil!`)
	}
}

func TestGetOmikujiDispatcherNilArgs(t *testing.T) {
	AssertPanic(t, "GetOmikujiDispatcher should have panicked!", func() {
		GetOmikujiDispatcher(nil, nil)
	})
}

func TestService_GetNextOmikujiWithDaikichi(t *testing.T) {
	currentTime := time.Now()
	currentMonth := currentTime.Month()
	currentDay := currentTime.Day()

	fromDate := PeriodicDate{Month: currentMonth, Day: currentDay}
	toDate := PeriodicDate{Month: currentMonth, Day: currentDay}
	pc := GetPeriodChecker(fromDate, toDate)

	dispatcher := GetOmikujiDispatcher(pc, &MockRandomizer{})

	omikuji := dispatcher.GetNextOmikuji()
	if omikuji.Text != "大吉" {
		t.Error(`Expected "大吉" omikuji! But was [`, omikuji.Text, `]`)
	}
}

func TestService_GetNextOmikujiWithNoDaikichi(t *testing.T) {
	futureTime := time.Now().AddDate(0, 1, 0)
	futureMonth := futureTime.Month()
	futureDay := futureTime.Day()

	fromDate := PeriodicDate{Month: futureMonth, Day: futureDay}
	toDate := PeriodicDate{Month: futureMonth, Day: futureDay}
	pc := GetPeriodChecker(fromDate, toDate)

	dispatcher := GetOmikujiDispatcher(pc, &MockRandomizer{})

	omikuji := dispatcher.GetNextOmikuji()
	if omikuji.Text != "吉" {
		t.Error(`Expected "吉" omikuji! But was [`, omikuji.Text, `]`)
	}
}

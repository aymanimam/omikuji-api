package omikuji

import (
	"fmt"
	"github.com/aymanimam/omikuji-api/errors"
	"time"
)

// PeriodChecker interfaces to check if a certain time instance lies in a certain period or not
type PeriodChecker interface {
	WithinThePeriod(t time.Time) bool
}

// Dispatcher interface to get the next random omikuji
type Dispatcher interface {
	GetNextOmikuji() Omikuji
}

// PeriodicDate a periodically repeated day every year
type PeriodicDate struct {
	Month time.Month
	Day   int
}

// period define a yearly periodic period
type period struct {
	From PeriodicDate
	To   PeriodicDate
}

// WithinThePeriod check if the given time in the defined period or not
func (p *period) WithinThePeriod(t time.Time) bool {
	layout := "2006-01-02"
	fromStr := fmt.Sprintf("%d-%02d-%02d", time.Now().Year(), p.From.Month, p.From.Day)
	toStr := fmt.Sprintf("%d-%02d-%02d", time.Now().Year(), p.To.Month, p.To.Day)

	fromTime, _ := time.Parse(layout, fromStr)
	toTime, _ := time.Parse(layout, toStr)

	fromTime = fromTime.Add(-24 * time.Hour)
	toTime = toTime.Add(24 * time.Hour)

	return t.After(fromTime) && t.Before(toTime)
}

// service Dispatcher service
type service struct {
	PeriodChecker PeriodChecker
	Randomizer    Randomizer
}

// GetNextOmikuji get the next random omikuji
func (s *service) GetNextOmikuji() Omikuji {
	if s.Randomizer == nil || s.PeriodChecker == nil {
		msg := fmt.Sprintf("One or more invalid arguments! [Randomizer: %v][PeriodChecker: %v]",
			s.Randomizer, s.PeriodChecker)
		errors.ThrowOmikujiException(msg, errors.OmikujiServiceErrorCode)
	}

	r := s.Randomizer
	currentTime := time.Now()

	if s.PeriodChecker.WithinThePeriod(currentTime) {
		return r.GetRandom(r.GetDaikichiMin(), r.GetMax())
	}
	return r.GetRandom(r.GetNoDaikichiMin(), r.GetMax())
}

// GetPeriodChecker Get PeriodChecker
func GetPeriodChecker(fromDate, toDate PeriodicDate) PeriodChecker {
	if fromDate.Month > toDate.Month {
		msg := fmt.Sprintf("Period checker inputs are invalid [fromDate: %v][toDate: %v]", fromDate, toDate)
		errors.ThrowOmikujiException(msg, errors.OmikujiServiceErrorCode)
	} else if fromDate.Month == toDate.Month {
		if fromDate.Day > toDate.Day {
			msg := fmt.Sprintf("Period checker inputs are invalid [fromDate: %v][toDate: %v]", fromDate, toDate)
			errors.ThrowOmikujiException(msg, errors.OmikujiServiceErrorCode)
		}
	}
	return &period{From: fromDate, To: toDate}
}

// GetOmikujiDispatcher Get Dispatcher
func GetOmikujiDispatcher(pc PeriodChecker, r Randomizer) Dispatcher {
	if pc == nil || r == nil {
		msg := fmt.Sprintf("Invalid arguments! [PeriodChecker: %v][Randomizer: %v]", pc, r)
		errors.ThrowOmikujiException(msg, errors.OmikujiServiceErrorCode)
	}
	return &service{
		pc,
		r,
	}
}

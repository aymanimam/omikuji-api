package http

import (
	"encoding/json"
	"github.com/aymanimam/omikuji-api/errors"
	"github.com/aymanimam/omikuji-api/middleware"
	"github.com/aymanimam/omikuji-api/omikuji"
	"log"
	"net/http"
	"time"
)

var omikujiDispatcher omikuji.Dispatcher

func initialize() error {
	// Initialize Daikichi period
	fromDate := omikuji.PeriodicDate{Month: time.January, Day: 1}
	toDate := omikuji.PeriodicDate{Month: time.January, Day: 3}
	omikujiRandomizer := omikuji.GetOmikujiRandomizer()
	periodChecker, err := omikuji.GetPeriodChecker(fromDate, toDate)
	if err != nil {
		return err
	}

	// Initialize omikuji service instance
	omikujiDispatcher, err = omikuji.GetOmikujiDispatcher(periodChecker, omikujiRandomizer)
	if err != nil {
		return err
	}

	return nil
}

func omikujiHandler(w http.ResponseWriter, r *http.Request) {
	// Get next random omikuji
	randOmikuji, err := omikujiDispatcher.GetNextOmikuji(time.Now())
	if err != nil {
		log.Printf("GetNextOmikuji error: %v", err)
		handleErrorResponse(w, err)
	} else {
		log.Printf("Omikuji: %v", randOmikuji)
		handleSuccessfulResponse(w, randOmikuji)
	}
}

func handleSuccessfulResponse(w http.ResponseWriter, o omikuji.Omikuji) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(o); err != nil {
		err := errors.NewOmikujiError(err.Error(), errors.OmikujiServerErrorCode)
		handleErrorResponse(w, err)
	}
}

func handleErrorResponse(w http.ResponseWriter, e error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	if jsonErr := json.NewEncoder(w).Encode(e); jsonErr != nil {
		log.Fatal(jsonErr)
	}
}

// StartServer start omikuji API server
func StartServer() {
	// Initialize
	err := initialize()
	if err != nil {
		log.Fatal(err)
	} else {
		// Centralized middleware for error handling
		r := middleware.NewRecovery()
		m := middleware.With(http.HandlerFunc(omikujiHandler), r)
		http.Handle("/omikuji", m)
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}
}

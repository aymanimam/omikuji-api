package server

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

func initialize() {
	// Initialize Daikichi period
	fromDate := omikuji.PeriodicDate{Month: time.January, Day: 1}
	toDate := omikuji.PeriodicDate{Month: time.January, Day: 3}
	periodChecker := omikuji.GetPeriodChecker(fromDate, toDate)
	omikujiRandomizer := omikuji.GetOmikujiRandomizer()

	// Initialize omikuji service instance
	omikujiDispatcher = omikuji.GetOmikujiDispatcher(periodChecker, omikujiRandomizer)
}

func omikujiHandler(w http.ResponseWriter, r *http.Request) {
	// Get next random omikuji
	randOmikuji := omikujiDispatcher.GetNextOmikuji()
	log.Printf("Omikuji: %v", randOmikuji)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(w).Encode(randOmikuji); err != nil {
		errors.ThrowOmikujiException(err.Error(), errors.OmikujiServerErrorCode)
	}
}

// StartServer start omikuji API server
func StartServer() {
	// Initialize
	initialize()

	// Centralized middleware for error handling
	r := middleware.NewRecovery()
	m := middleware.With(http.HandlerFunc(omikujiHandler), r)
	http.Handle("/omikuji", m)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

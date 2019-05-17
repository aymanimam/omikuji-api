package http

import (
	"github.com/aymanimam/omikuji-api/errors"
	"github.com/aymanimam/omikuji-api/middleware"
	"github.com/aymanimam/omikuji-api/omikuji"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
	"time"
)

var validResponseRegex = regexp.MustCompile(`^\\{"omikuji":"(大吉|吉|中吉|小吉|半吉|末吉|末小吉|凶|小凶|半凶|末凶|大凶)"\\}$`)

// Mock omikuji.Dispatcher
type MockPanicOmikujiDispatcher struct{}

func (omikujis *MockPanicOmikujiDispatcher) GetNextOmikuji(time time.Time) (omikuji.Omikuji, error) {
	msg := "MockPanicOmikujiDispatcher Error!"
	panic(errors.NewOmikujiError(msg, errors.OmikujiServiceErrorCode))
	return omikuji.Omikuji{}, nil
}

// ---------------

func TestOmikujiHandler(t *testing.T) {
	// Initialize
	initialize()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/omikuji", nil)
	omikujiHandler(w, r)
	rw := w.Result()

	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}

	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}

	if validResponseRegex.Match(b) {
		t.Fatalf("unexpected response: %s", string(b))
	}
}

func TestOmikujiHandlerErrorResponse(t *testing.T) {
	// Initialize
	omikujiDispatcher = &MockPanicOmikujiDispatcher{}

	rec := middleware.NewRecovery()
	mw := middleware.With(http.HandlerFunc(omikujiHandler), rec)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/omikuji", nil)

	mw.ServeHTTP(w, r)

	rw := w.Result()

	defer rw.Body.Close()

	if rw.StatusCode != http.StatusInternalServerError {
		t.Fatal("unexpected status code")
	}

	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}

	errStr := string(b)
	if !strings.Contains(errStr, "[MockPanicOmikujiDispatcher Error!") {
		t.Fatalf("unexpected response: %s", errStr)
	}
}

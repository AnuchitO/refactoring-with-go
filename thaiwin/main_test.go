package main

import (
	"bytes"
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckInHandler(t *testing.T) {
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(Check{ID: 1234, PlaceID: 4321})
	req := httptest.NewRequest("POST", "/checkin", payload)
	logger := zap.NewExample()
	l := logger.With(zap.String("middleware", "log"))
	r := req.WithContext(context.WithValue(req.Context(), "logger", l))

	w := httptest.NewRecorder()

	var fn = func(id, placeID int64) error {
		return nil
	}

	handler := &CheckIn{fn}

	handler.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Error("not ok")
	}
}

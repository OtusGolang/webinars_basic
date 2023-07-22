package service_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OtusTeam/Go-Basic/open_lessons/go_testing/service"
)

func TestVote(t *testing.T) {
	svc := service.NewService()

	rec := httptest.NewRecorder()
	// req := httptest.NewRequest(http.MethodPost, "/vote", bytes.NewBuffer([]byte(`{"passport":"pass", "candidate_id": 1}`)))

	req, err := http.NewRequest(http.MethodPost, "/vote", bytes.NewBuffer([]byte(`{"passport":"pass", "candidate_id": 1}`)))
	_ = err

	svc.ServeHTTP(rec, req)

	resp := rec.Result()
	_ = resp
}

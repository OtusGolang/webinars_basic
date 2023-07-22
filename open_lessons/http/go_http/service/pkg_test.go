package service_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OtusTeam/Go-Basic/open_lessons/go_http/service"
	"github.com/OtusTeam/Go-Basic/open_lessons/go_http/storage/inmem"
	"github.com/stretchr/testify/require"
)

func TestSubmitVote(t *testing.T) {
	cases := []struct {
		desc      string
		inMethod  string
		inPath    string
		inBody    []byte
		outStatus int
	}{
		{
			desc:      "case-1: test positive",
			inMethod:  http.MethodPost,
			inPath:    "/vote",
			inBody:    []byte(`{"passport":"pass", "candidate_id": 1}`),
			outStatus: http.StatusOK,
		},
		{
			desc:      "case-2: invalid method",
			inMethod:  http.MethodGet,
			inPath:    "/vote",
			inBody:    []byte(`{"passport":"pass", "candidate_id": 1}`),
			outStatus: http.StatusMethodNotAllowed,
		},
		{
			desc:      "case-3: corrupted JSON",
			inMethod:  http.MethodPost,
			inPath:    "/vote",
			inBody:    []byte(`{"passport":"pass", "candidate`),
			outStatus: http.StatusBadRequest,
		},
		{
			desc:      "case-4: validation failure",
			inMethod:  http.MethodPost,
			inPath:    "/vote",
			inBody:    []byte(`{"passport":"", "candidate_id": 1}`),
			outStatus: http.StatusBadRequest,
		},
	}

	srv := service.NewService(inmem.New())

	for i := range cases {
		tC := cases[i]

		rr := httptest.NewRecorder()
		req := httptest.NewRequest(tC.inMethod, tC.inPath, bytes.NewBuffer(tC.inBody))

		srv.SubmitVote(rr, req)

		resp := rr.Result()

		require.Equal(t, tC.outStatus, resp.StatusCode)
	}
}

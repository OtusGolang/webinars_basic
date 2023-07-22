package inmem_test

import (
	"testing"

	"github.com/OtusTeam/Go-Basic/open_lessons/go_http/storage/inmem"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	// Arrange
	storage := inmem.New()
	storage.Save(1) // 2
	storage.Save(1)
	storage.Save(2) // 1

	cases := []struct {
		desc          string
		inCandidateID uint32
		outVotes      uint32
		shouldFail    bool
		shouldFailIs  error
	}{
		{
			desc:          "case-01: get for existing",
			inCandidateID: 1,
			outVotes:      2,
			shouldFail:    false,
		},
		{
			desc:          "case-02: get for non existing",
			inCandidateID: 3,
			outVotes:      0,
			shouldFail:    true,
			shouldFailIs:  inmem.ErrNotFound,
		},
	}

	for i := range cases {
		tC := cases[i]

		res, err := storage.Get(tC.inCandidateID)

		if tC.shouldFail {
			require.Equal(t, uint32(0), res)
			require.ErrorIs(t, err, tC.shouldFailIs)
		} else {
			require.NoError(t, err)
			require.Equal(t, tC.outVotes, res)
		}
	}
}

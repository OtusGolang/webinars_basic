package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Service struct {
	Stats map[uint32]uint32
	sync.RWMutex
}

func NewService() *Service {
	return &Service{
		Stats: make(map[uint32]uint32),
	}
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/vote", s.SubmitVote)
	mux.HandleFunc("/stats", s.GetStats)

	mux.ServeHTTP(w, r)
}

func (s *Service) SubmitVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeErr(w, http.StatusMethodNotAllowed, fmt.Sprintf("method %s not not supported on uri %s", r.Method, r.URL.Path))

		return
	}

	buf := make([]byte, r.ContentLength)
	if _, err := r.Body.Read(buf); err != nil && err != io.EOF {
		writeErr(w, http.StatusBadRequest, err.Error())

		return
	}

	req := &VoteRequest{}
	if err := json.Unmarshal(buf, req); err != nil {
		writeErr(w, http.StatusBadRequest, err.Error())

		return
	}

	if req.Passport == "" || req.CandidateID == 0 {
		log.Printf("invalid arguments, skip vote")
		writeErr(w, http.StatusBadRequest, "invalid arguments")

		return
	}

	log.Printf("new vote receive (passport=%s, candidate_id=%d)",
		req.Passport, req.CandidateID)

	s.Lock()
	s.Stats[req.CandidateID]++
	s.Unlock()

	log.Print("vote accepted")

	w.WriteHeader(http.StatusOK)
}

func (s *Service) GetStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeErr(w, http.StatusMethodNotAllowed, fmt.Sprintf("method %s not not supported on uri %s", r.Method, r.URL.Path))
		return
	}

	args := r.URL.Query()
	id := args.Get("candidate_id")
	if len(id) > 0 {
		candidateId, err := strconv.Atoi(id)
		if err != nil {
			writeErr(w, http.StatusMethodNotAllowed, fmt.Sprintf("cant parse candidate_id, expect int, got: %s ", id))

			return
		}

		s.RLock()
		stat, ok := s.Stats[uint32(candidateId)]
		s.RUnlock()

		if !ok {
			writeErr(w, http.StatusBadRequest, fmt.Sprintf("candidate with id %d doasn't found", candidateId))

			return
		}

		resp := StatCandidateResponse{
			CandidateId: uint32(candidateId),
			Stat:        stat,
			Time:        time.Now(),
		}

		write(w, resp)

		return
	}

	s.RLock()
	stats := s.Stats
	s.RUnlock()

	resp := StatResponse{
		Records: stats,
		Time:    time.Now(),
	}

	write(w, resp)
}

func write(w http.ResponseWriter, resp any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	resBuf, err := json.Marshal(resp)
	if err != nil {
		log.Printf("response marshal error: %s", err)
	}

	if _, err = w.Write(resBuf); err != nil {
		log.Printf("response marshal error: %s", err)
	}
}

func writeErr(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	resBuf, err := json.Marshal(ErrResponse{msg})
	if err != nil {
		log.Printf("response marshal error: %s", err)
	}

	if _, err = w.Write(resBuf); err != nil {
		log.Printf("response marshal error: %s", err)
	}
}

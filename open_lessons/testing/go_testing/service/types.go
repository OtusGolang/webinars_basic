package service

import "time"

type VoteRequest struct {
	Passport    string `json:"passport,omitempty"`
	CandidateID uint32 `json:"candidate_id,omitempty"`
}

type ErrResponse struct {
	Error string `json:"message"`
}

type StatResponse struct {
	Records map[uint32]uint32 `json:"records,omitempty"`
	Time    time.Time         `json:"time,omitempty"`
}

type StatCandidateResponse struct {
	Time        time.Time `json:"time,omitempty"`
	CandidateId uint32    `json:"candidate_id,omitempty"`
	Stat        uint32    `json:"stat"`
}

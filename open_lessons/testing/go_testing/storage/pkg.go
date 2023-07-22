package storage

import "sync"

type InMemory struct {
	Stats map[uint32]uint32
	sync.RWMutex
}

func (s *InMemory) Save(candidateID uint32) {
	s.Lock()
	s.Stats[candidateID]++
	s.Unlock()
}

func (s *InMemory) GetByCandidateID(candidateID uint32) uint32 {
	s.RLock()
	stats := s.Stats[candidateID]
	s.RUnlock()

	return stats
}

func (s *InMemory) GetStats() map[uint32]uint32 {
	s.RLock()
	stats := s.Stats
	s.RUnlock()

	return stats
}

func New() *InMemory {
	return &InMemory{
		Stats: make(map[uint32]uint32),
	}
}

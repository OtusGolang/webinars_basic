package inmem

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("not found")

type InMem struct {
	m map[uint32]uint32
	sync.RWMutex
}

func (s *InMem) Save(cID uint32) error {
	s.RWMutex.Lock()
	s.m[cID]++
	s.RWMutex.Unlock()

	return nil
}

func (s *InMem) Get(cID uint32) (uint32, error) {
	s.RWMutex.RLock()
	stat, ok := s.m[cID]
	s.RWMutex.RUnlock()

	if !ok {
		return 0, ErrNotFound
	}

	return stat, nil
}

func (s *InMem) GetAll() (map[uint32]uint32, error) {
	s.RWMutex.RLock()
	stats := s.m
	s.RWMutex.RUnlock()

	return stats, nil
}

func New() *InMem {
	return &InMem{
		m: make(map[uint32]uint32),
	}
}

package main

import "sync"

type mutexService struct {
	mu sync.Mutex
	id uint64
}

func (s *mutexService) getNext() uint64 {
	s.mu.Lock()
	s.id++
	defer s.mu.Unlock()
	return s.id
}

package main

import "sync/atomic"

type atomicService struct {
	id uint64
}

func (s *atomicService) getNext() uint64 {
	atomic.AddUint64(&s.id, 1)
	return s.id
}

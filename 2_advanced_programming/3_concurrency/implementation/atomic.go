package main

import "sync/atomic"

type atomicService struct {
	id uint64
}

func (s *atomicService) getNext() uint64 {
	return atomic.AddUint64(&s.id, 1)
}

package main

type noSyncService struct {
	id uint64
}

func (s *noSyncService) getNext() uint64 {
	s.id++
	return s.id
}

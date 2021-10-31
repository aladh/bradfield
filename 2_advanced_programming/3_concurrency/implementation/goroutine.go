package main

type goroutineService struct {
	idChan chan uint64
}

func newGoroutineService() *goroutineService {
	service := &goroutineService{idChan: make(chan uint64)}
	go service.start()
	return service
}

func (s *goroutineService) getNext() uint64 {
	return <-s.idChan
}

func (s *goroutineService) start() {
	id := uint64(1)

	for {
		select {
		case s.idChan <- id:
			id += 1
		}
	}
}

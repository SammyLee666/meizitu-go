package scheduler

import (
	"Img/parser"
)

type SimpleScheduler struct {
	workerChan chan parser.Request
}

func (s *SimpleScheduler) Submit(r parser.Request) {
	go func() {
		s.workerChan <- r
	}()
}

func (s *SimpleScheduler) ConfigureWorkerChan(c chan parser.Request) {
	s.workerChan = c
}


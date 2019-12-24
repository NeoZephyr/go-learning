package scheduler

import "pain.com/go-learning/crawler/engine"

type DefaultScheduler struct {
	RequestPool chan engine.Request
}

func (s *DefaultScheduler) Submit(request engine.Request) {
	// 防止卡住
	//go func() {
	//	s.RequestPool <- request
	//}()
	s.RequestPool <- request
}

func (s *DefaultScheduler) ConfigureRequestPool(requestPool chan engine.Request) {
	s.RequestPool = requestPool
}




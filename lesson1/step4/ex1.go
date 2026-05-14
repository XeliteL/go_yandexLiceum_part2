package step4

import "sync"

type SafeMap struct {
	m  map[string]interface{}
	mu sync.RWMutex
}

func (s *SafeMap) Get(key string) interface{} {
	s.mu.RLock()
	data := s.m[key]
	s.mu.RUnlock()

	return data
}

func (s *SafeMap) Set(key string, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]interface{}),
	}
}

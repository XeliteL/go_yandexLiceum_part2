package step6

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

type SafeSLice struct {
	results []int
	mx      *sync.Mutex
}

func random() int {
	const max int = 100
	return rand.IntN(max)
}

func New() *SafeSLice {
	return &SafeSLice{
		results: []int{},
		mx:      &sync.Mutex{},
	}
}

func (s *SafeSLice) Append(item int) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.results = append(s.results, random())
}

func (s *SafeSLice) Get(index int) int {
	s.mx.Lock()
	defer s.mx.Unlock()

	return s.results[index]
}

func Theory1() {
	safeSLice := New()
	wg := &sync.WaitGroup{}
	const size int = 10

	wg.Add(size)
	for i := range size {
		go func(i int) {
			defer wg.Done()
			safeSLice.Append(random())
		}(i)
	}
	wg.Wait()

	for index := range size {
		fmt.Println(safeSLice.Get(index))
	}
}

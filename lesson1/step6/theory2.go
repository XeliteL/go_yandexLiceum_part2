package step6

import (
	"fmt"
	"sync"
	"time"
)

func Theory2() {
	var once sync.Once
	wg := &sync.WaitGroup{}
	initializeResources := func() {
		time.Sleep(time.Second)
		fmt.Println("xD")
	}

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(initializeResources)
		}()
	}
	wg.Wait()
}

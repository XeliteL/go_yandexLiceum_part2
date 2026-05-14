package step6

import (
	"sort"
	"strings"
	"sync"
)

func BestStudents(names []string) (string, error) {
	if len(names) == 0 {
		return "", nil
	}

	var mu sync.Mutex
	bestStudents := []string{}
	wg := &sync.WaitGroup{}
	var firstError error

	averageGrade, err := Average(names)
	if err != nil {
		return "", err
	}

	wg.Add(len(names))
	for _, name := range names {
		go func(studentName string) {
			defer wg.Done()

			studentGrade, err := getGrade(studentName)

			if firstError != nil {
				return
			}
			if err != nil {
				firstError = err
				return
			}

			if studentGrade > averageGrade {
				mu.Lock()
				defer mu.Unlock()
				bestStudents = append(bestStudents, studentName)
			}
		}(name)
	}
	wg.Wait()

	if firstError != nil {
		return "", firstError
	}

	mu.Lock()
	defer mu.Unlock()

	sort.Strings(bestStudents)

	result := strings.Join(bestStudents, ", ")
	return result, nil
}

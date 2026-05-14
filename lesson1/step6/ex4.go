package step6

import (
	"sync"
)

func CompateList(names []string) (map[string]string, error) {
	if len(names) == 0 {
		return nil, nil
	}

	var mu sync.Mutex
	categoryStudents := map[string]string{}
	wg := &sync.WaitGroup{}
	var firstError error

	averageGrade, err := Average(names)
	if err != nil {
		return nil, err
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

			mu.Lock()
			defer mu.Unlock()
			switch {
			case studentGrade > averageGrade:
				categoryStudents[studentName] = ">"
			case studentGrade < averageGrade:
				categoryStudents[studentName] = "<"
			default:
				categoryStudents[studentName] = "="
			}
		}(name)
	}
	wg.Wait()

	if firstError != nil {
		return nil, firstError
	}

	mu.Lock()
	defer mu.Unlock()

	return categoryStudents, nil
}

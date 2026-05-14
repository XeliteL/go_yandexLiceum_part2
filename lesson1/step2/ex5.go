package step2

import (
	"bufio"
	"errors"
	"os"
	"strings"
	"time"
)

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	f, err := os.OpenFile(inputFileName, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	result := []string{}

	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		logs := fileScanner.Text()
		log := strings.Split(logs, " ")
		t, err := time.Parse("02.01.2006", log[0])
		if err != nil {
			return nil, err
		}

		if !t.Before(start) && !t.After(end) {
			result = append(result, logs)
		}
	}

	if len(result) == 0 {
		return nil, errors.New("")
	}

	return result, err
}

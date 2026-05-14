package step6

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func Compare(name1, name2 string) (string, error) {
	grade1, err := getGrade(name1)
	if err != nil {
		return "", err
	}
	grade2, err := getGrade(name2)
	if err != nil {
		return "", err
	}

	switch {
	case grade1 > grade2:
		return ">", nil
	case grade1 < grade2:
		return "<", nil
	default:
		return "=", nil
	}
}

func getGrade(name string) (int, error) {
	url := "http://localhost:8082/mark?name=" + name
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	gradeString := strings.TrimSpace(string(body))
	grade, err := strconv.Atoi(gradeString)
	if err != nil {
		return 0, nil
	}

	return grade, nil
}

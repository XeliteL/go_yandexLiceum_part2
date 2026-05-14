package step7

import (
	"strings"
	"time"
)

func QuizRunner(questions, answers []string, answerCh chan string) int {
	count := 0
	for i := range questions {
		select {
		case userAnswer := <-answerCh:
			if strings.EqualFold(strings.TrimSpace(userAnswer), strings.TrimSpace(answers[i])) {
				count++
			}
		case <-time.After(time.Second):
		}
	}

	return count
}

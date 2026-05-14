package step2

import (
	"bufio"
	"os"
)

func LineByNum(inputFileName string, lineNum int) string {
	f, err := os.OpenFile(inputFileName, os.O_RDONLY, 0600)
	if err != nil {
		return ""
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	curLine := 0

	for fileScanner.Scan() {
		if curLine == lineNum {
			return fileScanner.Text()
		}
		curLine++
	}

	return ""
}

package step2

import (
	"os"
)

func ModifyFile(filename string, pos int, val string) {
	f, err := os.OpenFile(filename, os.O_RDWR, 0600)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = f.Seek(int64(pos), 0)
	if err != nil {
		return
	}

	f.WriteString(val)
}

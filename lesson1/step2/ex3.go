package step2

import (
	"io"
	"os"
)

func CopyFilePart(inputFilename, outFilename string, startpos int) error {
	iFile, err := os.OpenFile(inputFilename, os.O_RDONLY, 0600)
	if err != nil {
		return err
	}
	defer iFile.Close()

	oFile, err := os.Create(outFilename)
	if err != nil {
		return err
	}
	defer oFile.Close()

	offset := startpos
	_, err = iFile.Seek(int64(offset), 0)
	if err != nil {
		return err
	}

	buffer := make([]byte, 1024)

	for {
		n, err := iFile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		_, err = oFile.Write(buffer[:n])
		if err != nil {
			return err
		}
	}

	return nil
}

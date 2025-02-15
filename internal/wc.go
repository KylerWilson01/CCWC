package internal

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
)

const (
	// CouldNotOpenFileError is the error that happens when we fail to open a file
	CouldNotOpenFileError = "Could not open the given the file"
	// CouldNotGetStatsError is the error that happens when we fail to retrieve stats of a file
	CouldNotGetStatsError = "Could not get the stats of the given the file"
	// ReadError is the error that happens when the file cannot be read
	ReadError = "Could not read file"
)

// Record holds information about a file
type Record struct {
	ByteSize   int
	Lines      int
	Words      int
	Characters int
}

// AnalyzeFile gives the word count of a file with the given file path
func AnalyzeFile(f *os.File) (*Record, error) {
	fi, err := f.Stat()
	if err != nil {
		return nil, errors.New(CouldNotGetStatsError)
	}
	r := Record{
		ByteSize: int(fi.Size()),
	}

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		r.Lines++
		r.Words += len(bytes.Fields([]byte(line)))
		r.Characters += len(bytes.Runes([]byte(line)))

	}

	return &r, nil
}

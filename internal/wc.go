package internal

import (
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
	ByteSize int
	Lines    int
	Words    int
}

// AnalyzeFile gives the word count of a file with the given file path
func AnalyzeFile(fp string) (*Record, error) {
	f, err := os.Open(fp)
	if err != nil {
		return nil, errors.New(CouldNotOpenFileError)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return nil, errors.New(CouldNotGetStatsError)
	}
	r := Record{
		ByteSize: int(fi.Size()),
	}

	file := make([]byte, fi.Size())

	for {
		c, err := f.Read(file)
		r.Lines += bytes.Count(file[:c], []byte{'\n'})
		r.Words += len(bytes.Fields(file[:c]))

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.New(ReadError)
		}

	}

	return &r, nil
}

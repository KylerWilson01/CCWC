package internal

import (
	"errors"
	"os"
)

const (
	// CouldNotOpenFileError is the error that happens when we fail to open a file
	CouldNotOpenFileError = "Could not open the given the file"
	// CouldNotGetStatsError is the error that happens when we fail to retrieve stats of a file
	CouldNotGetStatsError = "Could not get the stats of the given the file"
)

// Record holds information about a file
type Record struct {
	ByteSize int
}

// AnalyzeFile gives the word count of a file with the given file path
func AnalyzeFile(fp string) (*Record, error) {
	f, err := os.Open(fp)
	if err != nil {
		return nil, errors.New(CouldNotOpenFileError)
	}

	fi, err := f.Stat()
	if err != nil {
		return nil, errors.New(CouldNotGetStatsError)
	}

	r := Record{
		ByteSize: int(fi.Size()),
	}

	return &r, nil
}

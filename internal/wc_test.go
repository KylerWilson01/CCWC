package internal_test

import (
	"errors"
	"os"
	"testing"

	"github.com/KylerWilson01/CCWC/internal"
)

func TestByteSize(t *testing.T) {
	tests := []struct {
		name, filePath string
		want           int
	}{
		{"Correct byte size", "testdata/test.txt", 342190},
	}

	for _, test := range tests {
		f, err := os.Open(test.filePath)
		if err != nil {
			panic(errors.New(internal.CouldNotOpenFileError))
		}
		defer f.Close()
		record, err := internal.AnalyzeFile(f)
		if err != nil {
			t.Errorf("\nCould not open %s\n", test.filePath)
		}

		if record.ByteSize != test.want {
			t.Errorf("\nWant: %d\nGot: %d", test.want, record.ByteSize)
		}
	}
}

func TestLineCount(t *testing.T) {
	tests := []struct {
		name, filePath string
		want           int
	}{
		{"Correct line count", "testdata/test.txt", 7145},
	}

	for _, test := range tests {
		f, err := os.Open(test.filePath)
		if err != nil {
			panic(errors.New(internal.CouldNotOpenFileError))
		}
		defer f.Close()
		record, err := internal.AnalyzeFile(f)
		if err != nil {
			t.Errorf("\nCould not open %s\n", test.filePath)
		}

		if record.Lines != test.want {
			t.Errorf("\nWant: %d\nGot: %d", test.want, record.Lines)
		}
	}
}

func TestWordCount(t *testing.T) {
	tests := []struct {
		name, filePath string
		want           int
	}{
		{"Correct word count", "testdata/test.txt", 58164},
	}

	for _, test := range tests {
		f, err := os.Open(test.filePath)
		if err != nil {
			panic(errors.New(internal.CouldNotOpenFileError))
		}
		defer f.Close()
		record, err := internal.AnalyzeFile(f)
		if err != nil {
			t.Errorf("\nCould not open %s\n", test.filePath)
		}

		if record.Words != test.want {
			t.Errorf("\nWant: %d\nGot: %d", test.want, record.Words)
		}
	}
}

func TestCharacterCount(t *testing.T) {
	tests := []struct {
		name, filePath string
		want           int
	}{
		{"Correct word count", "testdata/test.txt", 339292},
	}

	for _, test := range tests {
		f, err := os.Open(test.filePath)
		if err != nil {
			panic(errors.New(internal.CouldNotOpenFileError))
		}
		defer f.Close()
		record, err := internal.AnalyzeFile(f)
		if err != nil {
			t.Errorf("\nCould not open %s\n", test.filePath)
		}

		if record.Characters != test.want {
			t.Errorf("\nWant: %d\nGot: %d", test.want, record.Characters)
		}
	}
}

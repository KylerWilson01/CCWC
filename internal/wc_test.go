package internal_test

import (
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
		record, err := internal.AnalyzeFile(test.filePath)
		if err != nil {
			t.Errorf("\nCould not open %s\n", test.filePath)
		}

		if record.ByteSize != test.want {
			t.Errorf("\nWant: %d\nGot: %d", test.want, record.ByteSize)
		}
	}
}

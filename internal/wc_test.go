package internal_test

import (
	"testing"

	"github.com/KylerWilson01/CCWC/internal"
)

func TestWordCount(t *testing.T) {
	length, err := internal.WordCount("testdata/test.txt")
	if err != nil {
		t.Fatal()
	}

	if length != 342190 {
		t.Fail()
	}
}

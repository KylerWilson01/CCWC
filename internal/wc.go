package internal

import (
	"log"
	"os"
)

// WordCount gives the word count of a file with the given file path
func WordCount(fp string) (int, error) {
	f, err := os.Open(fp)
	if err != nil {
		log.Fatalf("Could not open the given file of %s", fp)
	}

	fi, err := f.Stat()
	if err != nil {
		log.Fatalf("Could not open the given file of %s", fp)
	}

	return int(fi.Size()), nil
}

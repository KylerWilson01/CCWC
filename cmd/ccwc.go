package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/KylerWilson01/CCWC/internal"
)

var (
	printBytes      bool
	printLines      bool
	printWords      bool
	printCharacters bool
)

func main() {
	flag.BoolVar(&printBytes, "c", false, "used to print the amount of bytes of a file")
	flag.BoolVar(&printLines, "l", false, "used to print the amount of lines in a file")
	flag.BoolVar(&printWords, "w", false, "used to print the amount of words in a file")
	flag.BoolVar(&printCharacters, "m", false, "used to print the amount characters in a file")
	flag.Parse()

	if !printBytes && !printLines && !printWords && !printCharacters {
		printBytes = true
		printLines = true
		printWords = true
	}

	fp := flag.Args()
	if len(fp) == 0 {
		fp = append(fp, "-")
	}

	for _, file := range fp {
		var f *os.File
		var err error

		if file == "-" {
			f = os.Stdin
		} else {
			f, err = os.Open(file)
			if err != nil {
				panic(errors.New(internal.CouldNotOpenFileError))
			}
			defer f.Close()
		}

		r, err := internal.AnalyzeFile(f)
		if err != nil {
			panic(err)
		}

		out := ""

		if printBytes {
			out = fmt.Sprintf("%s\t%d", out, r.ByteSize)
		}

		if printLines {
			out = fmt.Sprintf("%s\t%d", out, r.Lines)
		}

		if printWords {
			out = fmt.Sprintf("%s\t%d", out, r.Words)
		}

		if printCharacters {
			out = fmt.Sprintf("%s\t%d", out, r.Characters)
		}

		out = fmt.Sprintf("%s\t%s", out, f.Name())
		fmt.Println(out)
	}
}

package inpututil

import (
	"bufio"
	"log"
	"os"
)

type Seq[T any] func(yield func(T) bool)

func FileLines(filePath string) Seq[string] {
	return func(yield func(string) bool) {
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("failed to open file: %v", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				return
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("error reading file: %v", err)
		}
	}
}

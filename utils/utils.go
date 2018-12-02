package utils

import (
	"bufio"
	"log"
	"os"
	"time"
)

// ReadLines reads a whole file into memory
// and returns a slice of its lines.
// Comes from https://stackoverflow.com/a/18479916/9823697
func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// TimeTaken prints the time since "t".
// Should be called like "defer timeTaken(time.Now())"
func TimeTaken(t time.Time) {
	elapsed := time.Since(t)
	log.Printf("took %s\n", elapsed)
}

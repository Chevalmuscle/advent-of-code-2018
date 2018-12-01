package utils

import (
	"bufio"
	"os"
)

// ReadLines reads a whole file into memory
// and returns a slice of its lines.
// Comes from https://stackoverflow.com/a/18479916/9823697
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

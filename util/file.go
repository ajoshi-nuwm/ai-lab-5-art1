/*
Package util contains file utils for interacting with file system
 */
package util

import (
	"bufio"
	"os"
)

// Returns array of rows from file ot error in case of exception
func ReadFromFile(filename string) ([]string, error) {
	var read []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		read = append(read, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return read, nil
}

package cut

import (
	"cut/internal/key"
	"bufio"
	"os"
	"fmt"
)

func GetLines(key *key.Key) ([]string, error) {
	fileName := os.Args[len(os.Args) - 1]

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Error wheh open file: %v", err)
	}

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}


func Cut(key *key.Key) {

}
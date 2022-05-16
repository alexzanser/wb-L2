package cut

import (
	"bufio"
	"cut/internal/key"
	"fmt"
	"os"
	"strings"
)

type cut struct {
	key *key.Key
}

//New returns instance of type cut
func New(key *key.Key) *cut {
	return &cut{key: key}
}

//GetLines return slice of lines from file
func (*cut) GetLines() ([]string, error) {
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

//GetFields return list of specified fields
func (c *cut) GetFields() []string {
	index := make([]string, 0)

	for _, val := range c.key.Fields {
		index = append(index, strings.Split(val, "-")...)
	}

	return index
}

func (c *cut) Cut() {
	// lines, err := GetLines(key)
	// if err != nil {
	// 	return nil, fmt.Errorf("Error when get lines: %v", err)
	// }

	fmt.Println(c.GetFields())
}
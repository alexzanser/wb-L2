package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

type Key struct {
	after      int
	before     int
	context    int
	count      int
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	pattern    string
	filename   string
}

func InitKeys(rootCmd *cobra.Command, key *Key) {
	rootCmd.Flags().IntVarP(&key.after, "after-context", "A", 0, "after-context")
	rootCmd.Flags().IntVarP(&key.before, "before-context", "B", 0, "before-context")
	rootCmd.Flags().IntVarP(&key.context, "context", "C", 0, "context")
	rootCmd.Flags().IntVarP(&key.count, "count", "c", 0, "count")
	rootCmd.Flags().BoolVarP(&key.ignoreCase, "ignore-case", "i", false, "ignore-case")
	rootCmd.Flags().BoolVarP(&key.invert, "invert-match", "v", false, "invert-match")
	rootCmd.Flags().BoolVarP(&key.fixed, "fixed-strings", "F", false, "fixed-strings")
	rootCmd.Flags().BoolVarP(&key.lineNum, "line-number", "n", false, "line-number")
}

func GetLines(key *Key) ([]string, error) {

	var fileName string

	if len(os.Args) >= 3 {
		fileName = os.Args[2]
	} else {
		fileName = os.Stdin.Name()
	}

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

func Compare(key *Key, line, pattern string) bool {
	if key.ignoreCase {
		line, pattern = strings.ToLower(line), strings.ToLower(pattern)
	}

	if key.fixed {
		return line == pattern
	}

	ret, _ := regexp.MatchString(pattern, line)
	return ret
}

func Contains(lines []int, idx int) bool {
	for _, a := range lines {
		if a == idx {
			return true
		}
	}
	return false
}

type group struct {
	index []int
	left  int
	right int
}

func GetBorders(key *Key, idx int, len int) (int, int) {
	left := idx - key.before
	right := idx + key.after

	if left < 0 {
		left = 0
	}
	if right >= len {
		right = len - 1
	}

	return left, right
}

func UniquePrint(prevGroup, group *group, lines []string, idx int) {
	if len(group.index) == 0 && idx != 0 {
		if Contains(prevGroup.index, group.left) == false {
			fmt.Println("--")
		}
	}
	if Contains(prevGroup.index, idx) == false {
		fmt.Println(lines[idx])
		group.index = append(group.index, idx)
	}
}

func Grep(key *Key) ([]string, error) {
	if key.count > 0 {
		key.after = key.count
	}

	if key.context > 0 {
		if key.before == 0 {
			key.before = key.context
		}
		if key.after == 0 {
			key.after = key.context
		}
	}

	lines, err := GetLines(key)
	if err != nil {
		return nil, fmt.Errorf("Error when get lines: %v", err)
	}

	pattern := os.Args[1]
	prevGroup := &group{}
	gr := &group{}
	for i, line := range lines {
		if Compare(key, line, pattern) != key.invert {
			left, right := GetBorders(key, i, len(lines))
			gr = &group{left: left, right: right}
			for j := left; j <= right; j++ {
				UniquePrint(prevGroup, gr, lines, j)
			}
		}
		prevGroup.index = append(prevGroup.index, gr.index...)
		prevGroup.left = gr.left
		prevGroup.right = gr.right
	}
	return lines, nil
}

func main() {
	cmd := &cobra.Command{}
	key := &Key{}

	InitKeys(cmd, key)

	if err := cmd.Execute(); err != nil || len(os.Args) < 2 {
		log.Fatal(fmt.Errorf("required argument missing: %v", err))
	}

	Grep(key)
}

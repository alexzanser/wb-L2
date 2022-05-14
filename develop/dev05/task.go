package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"github.com/spf13/cobra"
)

type Key struct {
	after				int
	before				int
	context             int
	count				int
	ignoreCase			bool
	invert				bool
	fixed               bool
	lineNum				bool
	pattern				string
	filename			string
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

	file , err := os.Open(fileName)
	if 	err != nil {
		return nil, fmt.Errorf("Error wheh open file: %v", err)
	}

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil

}

func Grep(key *Key) ([]string, error) {
	if key.context > 0 {
		key.before = key.context
		key.after = key.context
	}

	// pattern := os.Args[1]

	lines, err := GetLines(key)
	if  err != nil {
		return nil, fmt.Errorf("Error when get lines: %v", err)
	}

	return lines, nil
}

func main() {
	cmd := &cobra.Command{}
	key := &Key{}

	InitKeys(cmd, key)

	if err := cmd.Execute(); err != nil || len(os.Args) < 2{
		log.Fatal(fmt.Errorf("required argument missing: %v", err))
	}

	Grep(key)
}

package main

import (
	"fmt"
	"log"
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
	diapason			
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

func Grep(key *Key) []string {
	
}

func main() {
	cmd := &cobra.Command{}
	key := &Key{}

	InitKeys(cmd, key)

	if err := cmd.Execute(); err != nil {
		log.Fatal(fmt.Errorf("required argument missing: %v", err))
	}

	Grep(key)
}

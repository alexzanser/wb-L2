package main

import (
	"fmt"
	"os"
	"log"
	"sort/internal/sort"
	"github.com/spf13/cobra"
	linesmodule "sort/internal/lines"
)


func main() {
	cmd := &cobra.Command{}
	key := &sort.Key{}

	sort.InitKeys(cmd, key)
	
	if err := cmd.Execute(); err != nil || len(os.Args) < 2 {
		log.Fatal(fmt.Errorf("required argument missing: %v", err))
	}
	
	lines, err := linesmodule.GetLines(os.Args[1])
	if err != nil {
		log.Fatal(fmt.Errorf("error when getting lines: %v", err))
	}

	if err := sort.CheckArguments(lines, key); err != nil {
		log.Fatal(fmt.Errorf("invalid arguments: %v", err))
	}

	sort.Sort(lines, key)

	for _, line := range lines {
		fmt.Println(line[1])
	}
}
package main

import (
	"fmt"
	"os"
	"log"
	"sort/internal/sort"
	keypackage "sort/internal/key"
	"github.com/spf13/cobra"
)


func main() {
	cmd := &cobra.Command{}
	key := keypackage.New()
	key.SetKeys(cmd)
	
	if err := cmd.Execute(); err != nil || len(os.Args) < 2 {
		log.Fatal(fmt.Errorf("required argument missing: %v", err))
	}
	
	lines := sort.Sort(key)
	for _, line := range lines {
		fmt.Println(line[1])
	}
}

package main

import (
	"fmt"
	key "cut/internal/key"
	cut "cut/internal/cut"
	"log"
	"os"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{}
	key := key.New()
	cut := cut.New(key)
	key.SetKeys(cmd)

	if err := cmd.Execute(); err != nil || len(os.Args) < 3 {
		log.Fatal(fmt.Errorf("required argument missing: %v", err))
	}
	
	lines, _ := cut.Cut()
	for _, line := range lines {
		fmt.Println(line)
	}
}

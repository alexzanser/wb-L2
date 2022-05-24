package main

import (
	"fmt"
	"log"
	"os"
	"github.com/spf13/cobra"
	key "grep/internal/key"
	grep "grep/internal/grep"
)

func main() {
	cmd := &cobra.Command{}
	key := key.New()

	key.SetKeys(cmd)

	if err := cmd.Execute(); err != nil || len(os.Args) < 2 {
		log.Fatal(fmt.Errorf("required argument missing: %v", err))
	}

	out, _ := grep.Grep(key)
	for _, l := range out {
		fmt.Println(l)
	}
}

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
	key.SetKeys(cmd)

	if err := cmd.Execute(); err != nil || len(os.Args) < 3 {
		log.Fatal(fmt.Errorf("required argument missing: %v", err))
	}

	cut := cut.New(key)
	cut.Cut()
}

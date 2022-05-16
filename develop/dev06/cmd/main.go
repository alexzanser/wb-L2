package main

import (
	"fmt"
	keymodule "cut/internal/key"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{}
	key := &keymodule.Key{}

	keymodule.InitKeys(cmd, key)

	if err := cmd.Execute(); err != nil || len(os.Args) < 2 {
		log.Fatal(fmt.Errorf("required argument missing: %v", err))
	}

	fmt.Println(key)
}
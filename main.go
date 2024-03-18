package main

import (
	"coin/cmd"
	"fmt"
	"os"
)

func main() {
	if err := cmd.RootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"github.com/mengfengkong/coin/cmd"
	"os"
)

func main() {
	if err := cmd.RootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

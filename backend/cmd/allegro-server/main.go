package main

import (
	"fmt"
	"os"

	"github.com/hoomanist/allegro-server/pkg/cli"
)

func main() {
	if err := cli.Root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

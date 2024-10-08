package main

import (
	"fmt"
	"os"

	"github.com/hoomanist/allegro-server/pkg/cli"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	if err := cli.Root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

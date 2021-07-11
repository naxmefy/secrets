package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "secrets",
		Usage: "secrets <command> [options]",
		Action: func(c *cli.Context) error {
			fmt.Println("boom!")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

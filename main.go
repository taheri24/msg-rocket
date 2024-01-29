package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"taheri24.ir/msg-rocket/server"
)

func main() {
	cliApp := cli.App{
		Commands: []*cli.Command{server.Cmd},
	}
	if err := cliApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

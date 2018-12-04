package main

import (
	"log"
	"os"

	"github.com/sizer/gorfc/fetch"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gorfc"
	app.Usage = "call gorfc with rfc no."
	app.Action = func(c *cli.Context) error {
		fetch.RfcDetail(c.Args().Get(0))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Army Knife"
	app.Usage = "Random day-to-day utilities"
	app.Action = func(c *cli.Context) {
		if !c.Args().Present() {
			fmt.Print("You must pass in at least one argument")
			return
		}
		if c.Args()[0] == "db" {
			ruby := RubyHelper{}
			ruby.databaseyml()
		}
	}
	app.Run(os.Args)
}

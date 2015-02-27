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
			fmt.Print("Current supported arguments are:  db, rmdd")
			return
		}
		if c.Args()[0] == "web" {
			web := WebHelper{}
			result := make(chan string, 1)
			result <- web.downloadJson(c.Args().Get(1))
			fmt.Println(<-result)
		}
		if c.Args()[0] == "db" {
			ruby := RubyHelper{}
			ruby.Databaseyml()
		}

		if c.Args()[0] == "var" {
			fle := NewFileHelper()
			fle.printPath(c.Args().Get(1))
		}

		if c.Args()[0] == "rmdd" {
			xc := NewXcodeHelper()
			xc.RemoveDerivedData()
		}
		if c.Args()[0] == "files" && c.Args()[1] == "-p" {
			fle := NewFileHelper()
			fle.prependValue = c.Args().Get(2)
			fle.listAllFiles(c.Args().Get(3))
		}
		if c.Args()[0] == "files" {
			fle := NewFileHelper()
			fle.listAllFiles(c.Args().Get(1))
		}

		if c.Args()[0] == "dirs" && c.Args()[1] == "-p" {
			fle := NewFileHelper()
			fle.prependValue = c.Args().Get(2)
			fle.listAllDirectories(c.Args().Get(3))
		}
		if c.Args()[0] == "dirs" {
			fle := NewFileHelper()
			fle.listAllDirectories(c.Args().Get(1))
		}
	}
	app.Run(os.Args)
}

package main

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func getCommands() []cli.Command {
	return []cli.Command{
		{
			Name:      "removederiveddata",
			ShortName: "rmdd",
			Usage:     "Removes derived data from Xcode",
			Action: func(c *cli.Context) {
				xc := NewXcodeHelper()
				xc.RemoveDerivedData()
			},
		},
		{
			Name:      "web",
			ShortName: "w",
			Usage:     "Download data using GET from webservice",
			Action: func(c *cli.Context) {
				web := WebHelper{}
				result := make(chan string, 1)
				result <- web.downloadJson(c.Args().First())
				fmt.Println(<-result)
			},
		},
		{
			Name:      "databaseyaml",
			ShortName: "db",
			Usage:     "Get the current dev database name from Rails",
			Action: func(c *cli.Context) {
				ruby := RubyHelper{}
				ruby.Databaseyml()
			},
		},
		{
			Name:      "var",
			ShortName: "v",
			Usage:     "Get the value of an evironment variable",
			Action: func(c *cli.Context) {
				fle := NewFileHelper()
				fle.printPath(c.Args().First())
			},
		},
		{
			Name:      "files",
			ShortName: "fs",
			Usage:     "List files in a directory [arguments...]",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "p",
					Value: "",
					Usage: "Value to prepend",
				},
			},
			Action: func(c *cli.Context) {
				fle := NewFileHelper()
				fle.prependValue = c.String("p")
				fmt.Println(fle.prependValue)
				fle.listAllFiles(c.Args().First())
			},
		},
		{
			Name:      "directories",
			ShortName: "d",
			Usage:     "List directories in a directory [arguments...]",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "p",
					Value: "",
					Usage: "Value to prepend",
				},
			},
			Action: func(c *cli.Context) {
				fle := NewFileHelper()
				fle.prependValue = c.String("p")
				fmt.Println(fle.prependValue)
				fle.listAllDirectories(c.Args().First())
			},
		},
	}
}

package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func databaseyml() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	check(err)

	data, err := ioutil.ReadFile(dir + "/config/database.yml")
	check(err)
	dat := string(data)

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(dat), &m)
	check(err)
	dev := m["development"].(map[interface{}]interface{})
	fmt.Print(dev["database"])
}

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
			databaseyml()
		}
	}
	app.Run(os.Args)
}

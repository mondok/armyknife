package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type RubyHelper struct{}

func (RubyHelper) databaseyml() {
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

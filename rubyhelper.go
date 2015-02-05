package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type RubyHelper struct{}

func (RubyHelper) Databaseyml() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if HasErr(err) {
		return
	}

	data, err := ioutil.ReadFile(dir + "/config/database.yml")
	if HasErr(err) {
		return
	}
	dat := string(data)

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(dat), &m)
	if HasErr(err) {
		return
	}
	dev := m["development"].(map[interface{}]interface{})
	fmt.Print(dev["database"])
}

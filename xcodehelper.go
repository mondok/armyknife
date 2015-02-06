package main

import (
	"log"
	"os/exec"
)

type XcodeHelper struct{}

func (XcodeHelper) RemoveDerivedData() {
	log.Print("Removing derived data")
	err := exec.Command("sh", "-c", "rm -rf ~/library/Developer/Xcode/DerivedData/*").Run()
	HasErr(err)
}

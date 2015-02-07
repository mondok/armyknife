package main

import (
	"log"
	"os/exec"
)

type XcodeHelper struct {
	DerivedDataPath string
}

func NewXcodeHelper() *XcodeHelper {
	xch := XcodeHelper{"rm -rf ~/library/Developer/Xcode/DerivedData/*"}
	return &xch
}

func (x XcodeHelper) RemoveDerivedData() {
	log.Print("Removing derived data from " + x.DerivedDataPath)
	err := exec.Command("sh", "-c", x.DerivedDataPath).Run()
	HasErr(err)
}

package main

import (
	"fmt"
	"os"
)

type FileHelper struct {
}

func NewFileHelper() *FileHelper {
	fh := FileHelper{}
	return &fh
}

func (f FileHelper) printPath(variableName string) {
	envVar := os.Getenv(variableName)
	fmt.Println(variableName + "=" + envVar)
}

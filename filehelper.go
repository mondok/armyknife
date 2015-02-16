package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileHelper struct {
	prependValue string
}

func NewFileHelper() *FileHelper {
	fh := FileHelper{}
	return &fh
}

func (f FileHelper) printPath(variableName string) {
	envVar := os.Getenv(variableName)
	fmt.Println(variableName + "=" + envVar)
}

func (f FileHelper) listAllFiles(path string) {
	err := filepath.Walk(path, f.visitFile)
	HasErr(err)
}

func (f FileHelper) listAllDirectories(path string) {
	err := filepath.Walk(path, f.visitDir)
	HasErr(err)
}

func (fh FileHelper) visitFile(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		return nil
	}
	fmt.Printf("%s %s\n", fh.prependValue, path)
	return nil
}

func (fh FileHelper) visitDir(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		fmt.Printf("%s %s\n", fh.prependValue, path)
	}
	return nil
}

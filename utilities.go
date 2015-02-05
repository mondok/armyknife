package main

import (
	"log"
)

func HasErr(e error) bool {
	if e != nil {
		log.Print(e)
		return true
	}
	return false
}

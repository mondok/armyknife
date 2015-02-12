package main

import (
	"log"
)

func HasErr(e error) bool {
	if e != nil {
		log.Print("Error")
		log.Print(e)
		return true
	}
	return false
}

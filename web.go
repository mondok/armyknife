package main

import (
	"io/ioutil"
	"net/http"
)

type WebHelper struct{}

func (WebHelper) downloadJson(url string) string {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if !HasErr(err) {
		return string(body)
	}
	return ""
}

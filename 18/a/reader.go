package main

import (
	"io/ioutil"
	"strings"
)

type FileReader struct {
	filename string
}

func (fr FileReader) ReadLines() []string {
	data, _ := ioutil.ReadFile(fr.filename)
	text := strings.TrimSpace(string(data))
	return strings.Split(text, "\n")
}

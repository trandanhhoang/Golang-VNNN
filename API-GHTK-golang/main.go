package main

import (
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.OpenFile("localfile.json", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString("new data that wasn't there originally\n"); err != nil {
		panic(err)
	}

	_, err = ioutil.ReadFile("localfile.json")
	if err != nil {
		panic(err)
	}

}

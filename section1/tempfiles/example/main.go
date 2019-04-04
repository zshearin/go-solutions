package main

import "github.com/zshearin/go-solutions/section1/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
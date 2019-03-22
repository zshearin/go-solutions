package main

import (
	"github.com/zshearin/go-solutions/section1/filedirs"
	"fmt"
)

func main() {
	if err := filedirs.Operate(); err != nil {
		panic(err)
	}

	fmt.Println("here")

	if err := filedirs.CapitalizerExample(); err != nil {
		panic(err)
	}

	
}


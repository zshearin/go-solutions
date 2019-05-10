package main

import (
	"fmt"

	"github.com/zshearin/go-solutions/section4/panic"
)

func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}

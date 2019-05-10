package main

import (
	"fmt"

	"github.com/zshearin/go-solutions/section4/errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()
	errwrap.Unwrap()
	fmt.Println()
	errwrap.StackTrace()
}

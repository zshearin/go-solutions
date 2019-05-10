package main

import "github.com/zshearin/go-solutions/section4/global"

func main() {
	if err := global.UseLog(); err != nil {
		panic(err)
	}
}

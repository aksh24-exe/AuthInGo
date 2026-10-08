package main

import (
	"GO-AUTH/app"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	cfg := app.Config{
		Addr: ":8080",
	}

	appliation := &app.Application{
		Config: cfg,
	}

	appliation.Run()
}
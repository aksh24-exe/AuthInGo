package main

import (
	"GO-AUTH/app"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	cfg := app.NewConfig(":8080")

	appliation := app.NewApplication(cfg)

	appliation.Run()
}
package main

import (
	"fmt"

	"github.com/LitusLuca/Taco/app"
)

func main() {
	fmt.Println("Hello World!")
	sandbox := app.App()
	sandbox.Run()
}

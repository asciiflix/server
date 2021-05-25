package main

import (
	"fmt"
	api "github.com/asciiflix/server/api"
)

func main() {
	fmt.Println("ASCIIflix Server")
	api.StartRouter()
}

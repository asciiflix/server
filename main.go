package main

import (
	"fmt"

	controller "github.com/asciiflix/server/controller"
	"github.com/asciiflix/server/database"
)

func main() {
	fmt.Println("ASCIIflix Server")
	database.StartDatabase()
	controller.StartRouter()
}

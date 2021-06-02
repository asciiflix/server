package main

import (
	"fmt"

	"github.com/asciiflix/server/config"
	"github.com/asciiflix/server/controller"
	"github.com/asciiflix/server/database"
)

var Version string

func main() {
	fmt.Println("ASCIIflix Server")
	config.GetConfig()
	db := database.StartDatabase()
	controller.StartRouter(db)
}

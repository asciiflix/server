package main

import (
	"fmt"

	"github.com/asciiflix/server/config"
	"github.com/asciiflix/server/controller"
	"github.com/asciiflix/server/database"
)

var Version string

func main() {
	fmt.Print("ASCIIflix Server ")
	fmt.Println(config.Version)
	config.GetConfig()
	database.StartDatabase()
	controller.StartRouter()
}

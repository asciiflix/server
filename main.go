package main

import (
	"fmt"

	"github.com/asciiflix/server/config"
	"github.com/asciiflix/server/controller"
	"github.com/asciiflix/server/database"
)

func main() {
	fmt.Println("ASCIIflix Server")
	config.GetConfig()
	db := database.StartDatabase()
	controller.StartRouter(db)
}

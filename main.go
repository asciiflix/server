package main

import (
	"fmt"

	api "github.com/asciiflix/server/api"
	"github.com/asciiflix/server/storage"
)

func main() {
	fmt.Println("ASCIIflix Server")
	db := storage.StartDatabase()
	api.StartRouter(db)
}

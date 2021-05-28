package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func StartRouter(db *gorm.DB) {
	r := mux.NewRouter()
	initHandler(r)
	fmt.Println("Starting API")
	log.Fatal(http.ListenAndServe(":8080", r))
}

package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartRouter() {
	r := mux.NewRouter()
	initHandler(r)
	fmt.Println("Stating API")
	log.Fatal(http.ListenAndServe(":8080", r))
}

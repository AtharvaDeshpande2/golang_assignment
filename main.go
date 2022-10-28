package main

import (
	"fmt"
	"github/atharvadeshpande/mongoapi/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Starting")
	log.Fatal(http.ListenAndServe(":8001", r))
	fmt.Println("Listening at port 8001.....")
}

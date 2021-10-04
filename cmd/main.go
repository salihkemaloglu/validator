package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/salihkemaloglu/validator/pkg/handler"
)

func main() {

	http.HandleFunc("/validate/vat-id/", handler.VatID)

	fmt.Printf("Starting server at port 8081\n")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

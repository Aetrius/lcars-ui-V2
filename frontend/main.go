package main

import (
	"log"
	"net/http"
)

type Aether struct {
	//exported field since it begins
	//with a capital letter
	CPU int
}

func main() {
	fs := http.FileServer(http.Dir("views"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":8000", nil)
}

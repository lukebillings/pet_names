package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	petname "github.com/dustinkirkland/golang-petname"
)

func Server() {
	log.Println("Starting Server")

	http.HandleFunc("/", home)
	http.HandleFunc("/names", petNames)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("ERRO: listen and serve: ", err)
	}
}

// most code in packages not in main

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "welcome")
}

func petNames(w http.ResponseWriter, req *http.Request) {
	words := 2
	separator := "-"

	wordsInt, err := strconv.Atoi(req.URL.Query().Get("words"))
	if err == nil {
		words = wordsInt
	}

	fmt.Fprint(w, petname.Generate(words, separator))
}

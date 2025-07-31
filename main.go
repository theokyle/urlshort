package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	mux := http.NewServeMux()
	godotenv.Load()
	port := ":" + os.Getenv("PORT")

	paths := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	mux.HandleFunc("/", hello)
	mapHandler := MapHandler(paths, mux)
	yamlHandler, err := YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		fmt.Printf("error handling yaml: %v", err)
	}

	fmt.Printf("Starting the sever on: %s\n", port)
	log.Fatal(http.ListenAndServe(port, yamlHandler))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

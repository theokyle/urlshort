package main

import (
	"flag"
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

	yamlFile := flag.String("yaml", "./paths.yaml", "a yaml file with a list of 'path:', 'url:' items")
	flag.Parse()

	yaml, err := os.ReadFile(*yamlFile)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
	}

	mux.HandleFunc("/", hello)
	yamlHandler, err := YAMLHandler(yaml, mux)
	if err != nil {
		fmt.Printf("error handling yaml: %v\n", err)
	}

	fmt.Printf("Starting the sever on: %s\n", port)
	log.Fatal(http.ListenAndServe(port, yamlHandler))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

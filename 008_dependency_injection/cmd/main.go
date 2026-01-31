package main

import (
	"dependency_injection"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello from dependency_injection main.go")

	dependency_injection.Greet(os.Stdout, "Hello Fikri")

	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependency_injection.MyGreeterHandler)))
}

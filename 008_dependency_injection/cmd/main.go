package main

import (
	"fmt"
	"github.com/adhitamafikri/go-with-test/dependency_injection"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello from dependency_injection main.go")

	dependency_injection.Greet(os.Stdout, "Hello Fikri")

	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependency_injection.MyGreeterHandler)))
}

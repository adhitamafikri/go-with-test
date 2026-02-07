package main

import (
	"fmt"
	"web_racer"
)

func main() {
	fmt.Println("Hello from 011_select")

	devToURL := "https://dev.to/"
	detikComURL := "https://www.detik.com/"
	result := web_racer.Racer(devToURL, detikComURL)

	fmt.Printf("The fastest URL is: %v", result)
}

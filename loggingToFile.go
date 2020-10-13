package main

import (
	"log"
	"os"
)

func main() {
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("./Error_logging/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Println("Hello world!")     //write output to file
	log.Println("hello world once again")
}

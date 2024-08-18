package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Please pass a text to encode.")
	}
	// asuming there's no way to pass less than 1 os.Args.
	//text := os.Args[1]
	//if (err != nil) {
	//	log.Fatal(err)
	//}
}
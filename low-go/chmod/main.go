package main

import (
	"log"
	"os"
)

func main() {

	if err := os.Chmod("hoge.txt", 0755); err != nil {
		log.Fatal(err)
	}
}

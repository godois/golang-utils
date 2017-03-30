package main

import (
	"os"
	"log"
	"fmt"
)

func main(){

	file, err := os.Open("/Users/marciogodoi/go/src/github.com/godois/golang-utils/file/products.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])

}

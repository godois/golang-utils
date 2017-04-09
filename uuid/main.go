package main

import (
	"os/exec"
	"log"
	"fmt"
)


type Person struct {
	name string
	age int
}


func main() {

	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
}

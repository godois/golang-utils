package main

import (
	"os/exec"
	"log"
	"fmt"
)

func main() {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
}

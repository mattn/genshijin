package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mattn/genshijin"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(genshijin.Shaberu(scanner.Text()))
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}

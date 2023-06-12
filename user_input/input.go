package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var print = fmt.Println

func main() {
	print("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, error := reader.ReadString('\n')
	if error == nil {
		print("You are welcome", name)
	} else {
		log.Fatal("an error occured: ", error)
	}
}

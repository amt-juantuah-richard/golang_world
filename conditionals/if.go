package main

import "fmt"

var pl = fmt.Println

func main() {
	theAge := 90
	if (theAge >= 1) && (theAge <= 18) {
		pl("too young to be married")
	} else if (theAge >= 18) && (theAge <= 25) {
		pl("should be in school for your age")
	} else if (theAge >= 25) && (theAge <= 35) {
		pl("get serious with your life")
	} else {
		pl("what are you doing here?")
	}
	pl("end of the life loop")
}

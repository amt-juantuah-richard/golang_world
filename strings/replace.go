package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main() {
	// replace a string with another string
	word := "The boy is wild and ready for success"
	newWord := strings.Replace(word, "wild and ready", "poised and prepared", -1)
	pl(newWord)

	// replace all white space characters with a "_"
	pl(strings.Replace(word, " ", "_", -1))

	// replace all "s" with 5
	pl(strings.Replace(word, "s", "5", -1))
}

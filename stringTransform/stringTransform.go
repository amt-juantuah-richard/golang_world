package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main() {
	var word = "Universal entity of ascii letters"

	// length of string
	pl("Length of the string is", len(word))

	// index of a character
	pl("Index of v in word variable is", strings.Index(word, "v"))
	pl("Index of i in word variable is", strings.Index(word, "i"))

	// trim spaces
	pl("Trim leading and trailing whitespaces:", strings.TrimSpace(word))
	pl("Trim leading and trailing unicode U:", strings.Trim(word, "U"))

	// split by a unicode character
	pl("Split on whitespaces:", strings.Split(word, " "))
	pl("Split on l:", strings.Split(word, "l"))

	// clone a string
	pl(strings.Clone(word))

	// compare two strings
	pl(strings.Compare(word, "Uax")) // result will be 0 if a == b, -1 if a < b, and +1 if a > b.

	// check if main string contains substring
	pl(strings.Contains(word, "U"))
	pl(strings.ContainsAny(word, "U"))
	pl(strings.ContainsRune(word, rune(5)))

	// count the occurence of a character
	pl(strings.Count(word, "l"))

	// seperate the string on a delimiter
	pl(strings.Cut(word, "of"))   // returns the string and first delimiter replaced with space
	pl(strings.Split(word, "of")) // returns the a new array with pieces of the seperated string

	pl(strings.Cut("a-b-c-d-e", "-"))
	pl(strings.Split("a-b-c-d-e", "-"))

	// join pieces of a string array together with a seperator string
	pl(strings.Join(strings.Split("a-b-c-d-e", "-"), ""))

	// tests whether the provided string starts with another string
	pl(strings.HasPrefix(word, "U"))

	// tests whether the provided string end with another string
	pl(strings.HasSuffix(word, "s"))

	// repeats a string a number of times
	pl(strings.Repeat("x\n", 10))
	pl(strings.Repeat("i\t", 2))

	// convert all string characters to lowercase
	pl(strings.ToLower(word))

	// convert all string to Title and Upper casing
	pl(strings.ToTitle(word))
	pl(strings.ToUpper(word))
}

package main

import (
	"fmt"
	"strings"
)

func GetCount(str string) (count int) {
	return strings.Count(str, "a") +
		strings.Count(str, "e") +
		strings.Count(str, "i") +
		strings.Count(str, "o") +
		strings.Count(str, "u")
}

func main() {
	fmt.Printf("input string: abracadabra, get vowel counts: %d\n", GetCount("abracadabra"))

}

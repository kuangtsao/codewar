package main

import (
	"fmt"
	"strings"
)

func GetCount(str string) (count int) {
	for _, c := range str {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}

func main() {
	fmt.Printf("input string: abracadabra, get vowel counts: %d\n", GetCount("abracadabra"))

}

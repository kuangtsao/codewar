package main

import (
	"fmt"
	"strings"
)

func GetCount(str string) (count int) {

	a_Ctr := strings.Count(str, "a")
	e_Ctr := strings.Count(str, "e")
	i_Ctr := strings.Count(str, "i")
	o_Ctr := strings.Count(str, "o")
	u_Ctr := strings.Count(str, "u")

	count = a_Ctr + e_Ctr + i_Ctr + o_Ctr + u_Ctr

	return count
}

func main() {
	fmt.Printf("input string: abracadabra, get vowel counts: %d\n", GetCount("abracadabra"))

}

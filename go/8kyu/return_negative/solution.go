package main

import "fmt"

func MakeNegative(x int) int {
	if x == 0 {
		return 0
	} else if x > 0 {
		return -x
	} else {
		return x
	}

}

func main() {
	fmt.Println("input  1,make it negative", MakeNegative(-1))
	fmt.Println("input -5,make it negative", MakeNegative(-5))
	fmt.Println("input  0,make it negative", MakeNegative(0))
}

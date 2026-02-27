package main

import (
	"fiveKyu"
	"fmt"
)


func main() {
	test1 := "Hello"
	test2 := "Emmanuel"
	test3 := "d🦊obabeedoo"
	fmt.Println(fiveKyu.FirstNonRepeating(test3))
	fmt.Println(fiveKyu.FirstNonRepeating(test2))
	fmt.Println(fiveKyu.FirstNonRepeating(test1))
}

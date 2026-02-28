package main

import (
	"fiveKyu"
	"fmt"
)

func main() {
	fmt.Println(fiveKyu.Alphanumeric("67fgfv"))
	fmt.Println(fiveKyu.Alphanumeric("hello world_"))
	fmt.Println(fiveKyu.Alphanumeric("\n\t\n"))
	fmt.Println(fiveKyu.Alphanumeric("     "))
}

// Expect(alphanumeric(".*?")).To(Equal(false))
// 		Expect(alphanumeric("a")).To(Equal(true))
// 		Expect(alphanumeric("Mazinkaiser")).To(Equal(true))
// 		Expect(alphanumeric("hello world_")).To(Equal(false))
// 		Expect(alphanumeric("PassW0rd")).To(Equal(true))
// 		Expect(alphanumeric("     ")).To(Equal(false))
// 		Expect(alphanumeric("")).To(Equal(false))
// 		Expect(alphanumeric("\n\t\n")).To(Equal(false))
// 		Expect(alphanumeric("ciao\n$$_")).To(Equal(false))
// 		Expect(alphanumeric("__ * __")).To(Equal(false))
// 		Expect(alphanumeric("&)))(((")).To(Equal(false))
// 		Expect(alphanumeric("43534h56jmTHHF3k")).To(Equal(true))
package main

import "fmt"

func main() {
	fmt.Println(beeramid(5000, 3))
}

func beeramid(refBonus, beerPrice int) int {
	count := 1
	comparer := 1
	numberOfCans := refBonus/beerPrice
	for i:= 3; i <= numberOfCans; i+=2 {
		comparer += i
		count++
		fmt.Println(numberOfCans)
		fmt.Println(comparer)
		if comparer == numberOfCans {
			break
		}
	}
	return count
}
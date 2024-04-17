package main

import (
	"fmt"
	"strconv"
)

func main() {
	sum := 1
	for i := 0; i < 100; i++ {
		checktext := ""
		if sum%3 == 0 {
			checktext += "Fizz"
		}
		if sum%5 == 0 {
			checktext += "Buzz"
		}
		if sum%3 != 0 && sum%5 != 0 {
			checktext += strconv.Itoa(sum)
		}
		fmt.Println(checktext)
		sum += 1
	}

	fmt.Println("TaDa!")
}

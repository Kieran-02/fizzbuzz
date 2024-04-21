package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i < 100; i++ {
		checktext := ""
		if i%3 == 0 || i%5 == 0 {
			if i%3 == 0 {
				checktext += "Fizz"
			}
			if i%5 == 0 {
				checktext += "Buzz"
			}
		} else {
			checktext += strconv.Itoa(i)
		}
		fmt.Println(checktext)
	}
	fmt.Println("TaDa!")
}

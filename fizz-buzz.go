package main

import "fmt"

func main() {
	//iterate through 1-100 via for loop
	i := 1
	for i <= 100 {
		fizzbuzz(i)
		i++
	}

}

func fizzbuzz(i int) {
	//switch case div by 3 = "fizz" div by 5 = "buzz" div by 3 & 5 = "fizzbuzz"
	switch {
	case (i%3 == 0) && (i%5 == 0):
		fmt.Print("(", i, ")")
		fmt.Println("FizzBuzz")
	case i%3 == 0:
		fmt.Print("(", i, ")")
		fmt.Println("Fizz")
	case i%5 == 0:
		fmt.Print("(", i, ")")
		fmt.Println("Buzz")
	default:
		fmt.Println(i)
	}
}

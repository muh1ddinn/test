package main

import "fmt"

func sum(l [10]int) (p, s int) {
	for i := 1; i < 10; i++ {
		if l[i]%2 == 0 {
			p = p * l[i]
		} else {
			s = s + l[i]
		}
	}
	return
}

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	p, s := sum(arr)
	fmt.Println(p) // prints the sum of even numbers
	fmt.Println(s) // prints the sum of odd numbers
}

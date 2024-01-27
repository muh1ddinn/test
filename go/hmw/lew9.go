package main

import "fmt"

func kk(k []int) (odd, even, mus, manf, nol []int) {
	for i := 0; i < len(k); i++ {
		if k[i]%2 == 0 && k[i] != 0 {

			even = append(even, k[i])
		} else {

			odd = append(odd, k[i])

		}
		if k[i] < 0 {

			manf = append(manf, k[i])
		} else if k[i] > 0 {
			mus = append(mus, k[i])

		} else if k[i] == 0 {
			nol = append(nol, k[i])
		} else {
			fmt.Print(" ")
		}

	}
	return odd, even, mus, manf, nol

}

func main() {
	var n int
	fmt.Println("Please enter the number of integers you want to input:")
	fmt.Scan(&n)

	k := make([]int, 0, n)

	fmt.Println("Please enter your numbers:")
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		k = append(k, num)
	}

	t, r, e, y, s := kk(k)

	fmt.Println("Even numbers:", y)
	fmt.Println("Odd numbers:", s)
	fmt.Println("Negative numbers:", r)
	fmt.Println("Positive numbers:", t)
	fmt.Println("Zero numbers:", e)
}

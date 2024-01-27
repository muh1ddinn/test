package main

import "fmt"

func classifyNumbers(k []int) (even, odd, manf, mus, nol []int) {
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
		}
	}
	return
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

	even, odd, manf, mus, nol := classifyNumbers(k)

	fmt.Println("Odd numbers:", odd)
	fmt.Println("Even numbers:", even)
	fmt.Println("Negative numbers:", manf)
	fmt.Println("Positive numbers:", mus)
	fmt.Println("nol:", nol)
}

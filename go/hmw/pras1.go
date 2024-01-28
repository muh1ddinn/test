package main

import "fmt"

func classifyNumbers(k []int) (ev, od, man, muss, noll int) {

	var even, odd, manf, mus, nol []int
	for i := 0; i < len(k); i++ {
		if k[i]%2 == 0 && k[i] != 0 {
			even = append(even, k[i])
			ev++
		} else if k[i]%2 != 0 && k[i] != 0 {
			odd = append(odd, k[i])
			od++
		}
		if k[i] < 0 {
			manf = append(manf, k[i])
			man++
		} else if k[i] > 0 {
			mus = append(mus, k[i])
			muss++
		} else if k[i] == 0 {
			nol = append(nol, k[i])
			noll++
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

	ev, od, man, muss, noll := classifyNumbers(k)

	fmt.Println("Odd numbers:", od)
	fmt.Println("Even numbers:", ev)
	fmt.Println("Negative numbers:", man)
	fmt.Println("Positive numbers:", muss)
	fmt.Println("nol:", noll)
}

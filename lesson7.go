package main

import "fmt"

func kk(a []int) (even []int, odd []int,musbat[]int,manfiy[]int) {
    for _, v := range a {
        if v%2 != 0 {
            odd = append(odd, v)
        } else {
            even = append(even, v)
        }   
        if v > 0 {

            musbat = append(musbat ,v)
        }else {
            manfiy = append(manfiy , v)
        }
    }
    return odd, even, musbat, manfiy 
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

    even, odd, musbat, manfiy := kk(k)

    fmt.Println("Even numbers:", even)
    fmt.Println("Odd numbers:", odd)
    fmt.Println("manfiy numbers:", manfiy)
    fmt.Println("musbat numbers:", musbat)

}

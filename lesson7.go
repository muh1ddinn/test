package main
import ("fmt")

func myFunction(a, b, c int) int {
    if a > b {
        if b > c {
            return b
        } else if a > c {
            return c
        } else {
            return a
        }
    } else {
        if a > c {
            return a
        } else if b > c {
            return c
        } else {
            return b
        }
    }
}





func main() {
  
    c:=myFunction(10,15,20)
  fmt.Println(c)
  
}
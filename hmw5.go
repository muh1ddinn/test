package main 

import "fmt"

/*
func cc(i,o int ) float64 {
    var p float64

     p=float64(i)/float64(o)
    
      return p
}



func main (){

  k:=cc(56,34)

  fmt.Println(k)
 
}*/





  func sum( y int ) int {
   
    var  u int 

    for i:=0;i<=y;i++{
    
      u=u+i

    } 
       return u

  }

func main(){

   c:=sum(7)

   fmt.Println(c)

}
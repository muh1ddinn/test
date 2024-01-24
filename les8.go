package main


import "fmt"



func sum(arr [] int)int{


	var c int

	for i:=2;i<len(arr)-6;i++{
     
		c=c+arr[i]

	}
    return c

}
 


func main(){


 
  arrr:=[10]int{1,-2,3,-4,5,-6,7,-8,9,-10}


 
   
  fmt.Println(sum(arrr[:]))



}
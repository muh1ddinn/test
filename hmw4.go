package main

import "fmt"


func kk(x,f,g,h int)bool{


   if (x+f+g+h)%2==0 {

	return true
   }else {

	return false
   }

}




func main(){
 
	c:=kk(3,5,2,2)
     
	fmt.Println(c)




}
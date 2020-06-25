package main

import "fmt"

func main(){

  if 9%2==0 {
    fmt.Println("9 is even")
  } else{
    fmt.Println("9 is odd")
  }

  if 10%5==0 {
    fmt.Println("10 is divisible by 5")
  }

  if num :=9; num<0 {
    fmt.Println(num, "is negative")
  } else if num < 10 {
    fmt.Println(num, "has 1 digit")
  } else {
    fmt.Println(num, "has multiple digits")
  }

}

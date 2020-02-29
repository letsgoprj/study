package main

import (
	"fmt"
	"unsafe"
)

func main() {
	ch := make(chan int)
	
	go func() {
		ch <- 10
	}()
	
	result := <-ch
	fmt.Println(result);
	
	num := 5
	pNum := &num
	
	fmt.Println("num : ", num)
	fmt.Println("pnum : ", pNum)
	fmt.Println("pnum : ", *pNum)
	
	*pNum++;
	fmt.Println("num : ", num)
	fmt.Println("pnum : ", pNum)
	fmt.Println("pnum : ", *pNum)
	
	var n1, n2, n3 int
	
	fmt.Print("Input num : ")
	fmt.Scanln(&n1, &n2, &n3)
	fmt.Println(n1, n2, n3)
	
	s1 := "12345678901234567890"
	d1 := 0.1
	b1 := false
	fmt.Printf("(%T)%d, (%T)%d, (%T)%d, (%T)%d \n", 
			   n1, unsafe.Sizeof(n1), s1, unsafe.Sizeof(s1), d1, unsafe.Sizeof(d1), b1, unsafe.Sizeof(b1))
	
	// var num1, num2, num3 int
	
	// fmt.Scanln(&num1, &num2, &num3)
	// c1 := float32(num1)
	// c2 := uint(num2)
	// c3 := int64(num3)
	
	// fmt.Printf("%T, %f \n", c1, c1)
	// fmt.Printf("%T, %d \n", c2, c2)
	// fmt.Printf("%T, %d \n", c3, c3)
	
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Printf("%v \n", arr)
}
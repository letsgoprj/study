package main

import "fmt"

func main() {
	
	//for_ex()	// for 
	//gugudan()	// 구구단
	//triangle()	// 이등변삼각형
}

func for_ex()() {
	var arr [5]int = [5]int{ 1, 2, 3, 4, 5 }
	
	for i, n := range arr {
		fmt.Printf("Arr[%d] = %d \n", i, n)
	}
	
	var m map[string]string = map[string]string {
		"A" : "AAA",
		"B" : "BBB",
	}
	
	for x, y := range m {
		fmt.Printf("x(%s) : y(%s) \n", x, y)
	}
}

func gugudan() {
	var dan int
	
	fmt.Print("Input Num : ")
	fmt.Scanln(&dan)
	
	for i := 1; i <= 9; i++ {
		fmt.Printf("%d X %d = %d\n", dan, i, dan * i)
	}
}

func triangle() {
	var i, j, n int 
	
	fmt.Print("Input Num : ")
	fmt.Scanln(&n)
	
	for i = 0; i < n; i++ {
		for j = 0; j < i; j++ {
			fmt.Printf("o ")
		}
		fmt.Printf("* \n")
	}
}

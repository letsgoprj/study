package main

import "fmt"

var global_a = 100;

func main() {
	i := 10		// 자동 타입 추론 (함수내에서만)
	var a int = 5;
	var b int;	// zero value 초기화, a = 0
	
	//var c int = 10; // 선언 후 쓰지 않으면 compile error
	const c int = 10; // 상수는 no error
	
	fmt.Printf("hello, Golang! \n")
	fmt.Printf("i = %d, a = %d \n", i, a)
	fmt.Print("b = ", b)
	fmt.Println("global_a = ", global_a)
	
	var d = 1.0;
	fmt.Printf("(%T) d = %f, %d \n", d, d, d)
	fmt.Printf("(%T) d = %f, %d \n", b, b, b)
	
	e, f, g := 1, "2", 3.0
	fmt.Printf("(%T, %T, %T) \n", e, f, g);
}

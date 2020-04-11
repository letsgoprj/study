package main

import "fmt"

func main() {
	//if_format()	
	//multiple()
	//diff_two_num()
	//calculator()
	//gugudan2()
	sum99()
}

func if_format() {
	var n int 
	
	fmt.Print("Input Num : ")
	fmt.Scan(&n)
	
	if n == 1 {
		fmt.Print("1\n")
	} else {
		fmt.Print("2\n")
	}
	
	if val := n * 2; val == 2 {
		fmt.Printf("val=%d, n=%d\n", val, n)
	}
}

func multiple() {
	
	for i := 1; i < 100; i++ {
		if i % 7 == 0 || i % 9 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("\n")
}

func diff_two_num() {
	var n1, n2, result int 
	
	fmt.Print("Input Num 1,2 : ")
	fmt.Scan(&n1, &n2)
	
	if n1 > n2 {
		result = n1 - n2;
	} else {
		result = n2 - n1;
	}
	
	fmt.Printf("%d\n", result)
}

func calculator() {
	var sel int 
	var num1, num2, result float64
	
	fmt.Print("Operator 1:+, 2:-, 3:x, 4:/ : ")
	fmt.Scan(&sel)
	
	fmt.Print("Input Num 1,2 : ")
	fmt.Scan(&num1, &num2)
	
	switch sel {
	case 1:
		result = num1 + num2
	case 2:
		result = num1 - num2
	case 3:
		result = num1 * num2
	case 4:
		result = num1 / num2
	default:
		fmt.Print("잘못된입력입니다.\n")
		return
	}
	
	fmt.Printf("%.1f\n", result)
}

func gugudan2() {
	
	for i := 2; i <= 9; i++ {
		if i % 2 == 0 {
			continue
		}
		for j := 1; j <= 9; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i * j)
			if i == j {
				break;
			}
		}
		fmt.Println()
	}
}

func sum99() {
	var result int 
	
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			if i == j {
				continue
			}
			result = (i * 10 + j) + (j * 10 + i);
			if (result == 99) {
				fmt.Printf("%02d + %02d = %d\n", i * 10 + j, j * 10 + i, result)
			}
		}
	}
}

package main

import "fmt"

func main() {
	//arr_func()
	//slice()
	//make_slice()
	//append_slice()
	//map_ex()
	//long_name()
	test_aver()
}

func arr_func() {
	arr := [...]int { 1, 2, 3 }
	fmt.Println(arr, len(arr))
}

func slice() {
	var a []int = []int { 1, 2, 3 }
	
	a[1] = 10
	
	fmt.Println(a)
	
	var b []int 
	
	if b == nil {
		fmt.Println("cap: ", cap(b), "len:", len(b), " nil slice")
	}
}

func make_slice() {
	s := make([]int, 0, 3)
	
	for i := 1; i <= 13; i++ {
		s = append(s, i)
		fmt.Println(len(s), cap(s))
	}
	
	fmt.Println(s)
}

func append_slice() {
	sliceA := []int { 1, 2, 3 }
	sliceB := []int { 4, 5, 6 }
	
	sliceA = append(sliceA, sliceB...)
	
	fmt.Println(sliceA)
	
	sliceC := make([]int, len(sliceA), cap(sliceA) * 2)
	
	copy(sliceC, sliceA)
	
	fmt.Println(sliceC)
	fmt.Println(len(sliceC), cap(sliceC))
	
	sliceD := sliceC[:3]
	//sliceD := make([]int, len(sliceC), cap(sliceC))
	//copy(sliceD, sliceC)
	fmt.Println(sliceD)
	
	sliceD[0] = -1
	fmt.Println(sliceD, sliceC, sliceA)
}

func map_ex() {
	var m = make(map[string]string)
	
	m["01"] = "AAA"
	m["02"] = "BBB"
	m["03"] = "CCC"
	
	fmt.Println(m)
	
	m["03"] = "CCC-1"
	
	fmt.Println(m)
	
	delete(m, "02")
	
	fmt.Println(m)
	
	m["04"] = "DDD"
	
	fmt.Println(m["04"])
	fmt.Println(m["05"])
	
	val, exist := m["03"]
	fmt.Println(val, exist)
	
	_, exist = m["04"]
	_, exist2 := m["05"]
	fmt.Println(exist, exist2)
	
	fmt.Println(m, len(m))
}

func long_name() {
	names := make([]string, 0, 1)
	var name string
	
	for {
		fmt.Print("Input Name: ")
		fmt.Scanln(&name)
		if name == "1" {
			break;
		}
		names = append(names, name)
	}
	fmt.Println(names, len(names[0]))
	
	var result string = names[0]
	
	for _, val := range names {
		if len(val) > len(result) {
			result = val
		}
	}
	
	fmt.Println(result, len(result))
}

func test_aver() {
	m := make(map[string]int)
	var avg float32
	
	for {
		var subject string
		var score int
		
		fmt.Print("Input Subject, Score: ")
		fmt.Scanln(&subject, &score)
		if subject == "0" {
			break;
		}
		m[subject] = score
	}
	fmt.Println(m)
	
	count := len(m)
	for key, val := range m {
		fmt.Println(key, val)
		avg += float32(val);
	}
	fmt.Printf("%.2f\n", avg / float32(count))
}
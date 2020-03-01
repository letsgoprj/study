package main

func main() {
	var a int
	var f float32 = 11.
	a = 10
	f = 12.0
	var i, j, k int = 1, 2, 3
	// 변수를 선언하면서 초기값을 지정하지 않으면, Go는 Zero Value를 기본적으로 할당한다. 즉, 숫자형에는 0, bool 타입에는 false, 그리고 string 형에는 "" (빈문자열)을 할당한다.

	const c int = 10 //const c = 10은 타입 추론 기능
	const s string = "Hi" //const s = "Hi"

	const (			//상수 괄호 안에 상수 나열하여 묶어서 사용할 수 있다.
		Visa = "Visa"
		Master = "MasterCard"
		Amex = "American Express"
	)

	const (
		Apple = iota //0, iota는 identifier
		Grape //1
		Orange //2
	)
	// 예약 키워드
	//break        default      func         interface    select
	//case         defer        go           map          struct
	//chan         else         goto         package      switch
	//const        fallthrough  if           range        type
	//continue     for          import       return       var
}

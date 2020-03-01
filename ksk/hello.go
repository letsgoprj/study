package main //일발적으로 패키지는 라이브러리로서 사용되지만, go compiler에 의해 특별하게 인식, 패키지명이 main인 경우 컴파일러는 해당 패키지를 공유 라이브러리가 아닌 실행 프로그램을 만든다.
import "fmt" //go의 표준 라이브러리 println()함수 호출 가능

func main() {
	fmt.Printf("Hello, world\n");
}

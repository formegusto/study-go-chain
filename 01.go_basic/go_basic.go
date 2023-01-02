package go_basic

import (
	"fmt"
	"strconv"
)

// short cut은 함수 밖에서 이용 불가능하다.
// globalName := "gusto"
// 무조건 long syntax로 이용
const globalName string = "gusto"

func Variable() {
	// type-safe
	// var name string = "forme"
	// error!
	// name = 12

	// short cut
	// type 추론
	name := "forme"
	// error!
	// error!
	// name = 12
	age := 28
	// := (only create), = (update)

	// const 는 상수로 변경 불가능하다.
	// globalName = "th"

	fmt.Println(name, age)	
}

func FuncExam() func(a, b int) (int, string) {
	fmt.Println("지금 당신이 하고 있는 것이 함수 만들기")

	plus := func(a int, b int) (int, string) {
		return a + b, strconv.Itoa(a + b)
	}
	return plus
}

func Sum(numbers ...int) int {
	// sum := 0
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func ShowLetter(str string) {
	for _, letter := range str {
		// byte 반환
		// fmt.Println(letter)

		// 해결방법
		// fmt.Println(string(letter))

		// fotmatting 예시
		fmt.Printf("%b\n",letter)
	}
}

func Fmt() {
	x := 405940594059
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%o\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%U\n", x)

	// Sprintf
	xAsBinary := fmt.Sprintf("%b\n", x)
	fmt.Println(x,xAsBinary)
}
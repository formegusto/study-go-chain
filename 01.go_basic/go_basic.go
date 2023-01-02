package go_basic

import (
	"fmt"
	"strconv"

	"github.com/formegusto/study-go-chain/01.go_basic/person"
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

// js와 python과 다르게 정해진 길이가 있다.
func Arrays() {
	foods := [3]string{"pizza", "hamburger", "potato"}
	// for _, food := range foods {
	// 	fmt.Println(food)
	// }
	// %v is default format
	fmt.Printf("%v\n", foods)
	for i:=0; i<len(foods);i++ {
		fmt.Println(foods[i])
	}

	// Slice
	names := []string{"kim", "lee", "park"}
	for _, name := range(names) {
		fmt.Println(name)
	}
	// names = append(names, "no")
	namesNew := append(names, "no")
	for _, name := range(namesNew) {
		fmt.Println(name)
	}
}

func Pointers() {
	a := 2
	// copy value 
	b := a 
	b = 3 
	fmt.Println(a, b) // 2, 3

	// copy position
	c := &a
	(*c) = 3
	fmt.Println(a, *c) // 3, 3
}

// class와 가장 유사한 개념
// type person struct {
// 	name 	string
// 	age 	int
// }

// // receiver
// func (instance person) sayHello() {
// 	fmt.Printf("Hello! My name is %s and I'm %d\n", instance.name, instance.age)
// }

// func Structs() {
// 	// forme := person{"gusto", 28}
// 	forme := person{name: "gusto", age: 28}
// 	fmt.Println(forme)

// 	forme.sayHello()
// }

func PointersAndStructs() {
	forme := person.Person{}
	fmt.Println(forme)

	forme.SetDetails("gusto", 28)
	fmt.Println(forme)
}
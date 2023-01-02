package person

type Person struct {
	name 	string
	age 	int
}

// Setting 때, Pointer의 개념이 필요해서 중요하다고 하신거다.
// 여기는 call by value 이기 때문
// 근데 structure는 그냥 박히는 것이 너무 신기하당.
func (p *Person) SetDetails(name string, age int) {
	p.name = name
	p.age = age
}
package person

type Person struct {
	name string
	age int
}

func (p *Person) SetName(name string) {
	p.name = name;
}

func (p *Person) SetAge(age int) {
	p.age = age;
}
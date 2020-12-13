package main

import (
	"fmt"
)

func main() {
	intPointerApp()
    structPointerApp()
	structMethodApp()
}

func intPointerApp() {
	fmt.Println()
	fmt.Println("=== int pointer app")

	num := 100
	p := &num

	fmt.Printf("num addr: %p, num value: %v, p value: %v\n", p, num, *p)

	*p = *p * 6

	fmt.Printf("num addr: %p, num value: %v, p value: %v\n", p, num, *p)

}

type Vertex struct {
	X int
	Y int
}

type Person struct {
	gender string
	age int
	name string
}

type Student struct {
	grade string
	Person
}

type Young interface {
	Name() string
	SetName(string)
}

type AdvanceInt int

func structPointerApp() {
	fmt.Println()
	fmt.Println("=== struct pointer app")

	v := Vertex{Y: 300}
	fmt.Printf("v vlaue: %v\n", v)

	p := &v
	p.X = 300
	fmt.Printf("v vlaue: %v\n", v)

}

func (ai AdvanceInt) Abs() int {
	if (ai > 0 ) {
		return int(ai)
	} else {
		return int(-1 * ai)
	}
}

func (person Person) String() string {
	return fmt.Sprintf("name: %v, age: %v, gender: %v",
		person.name, person.age, person.gender)
}

func (student Student) String() string {
	return fmt.Sprintf("person: (%v), grade: %v",
		student.Person, student.grade)
}

func (student Student) Name() string {
	return student.name
}

func (student *Student) SetName(name string) {
	student.name = name
}

func structMethodApp() {
	fmt.Println()
	fmt.Println("=== struct method app")

	ai := AdvanceInt(-900)
	fmt.Println(ai.Abs())

	person := Person{name: "jack", age: 18, gender: "male"}
	fmt.Println(person)

	student := Student{grade: "大一", Person: person}
	fmt.Println(student)

	fmt.Println(student.Name())
	student.SetName("pain")
	fmt.Println(student.Name())
	student.name = "page"

	var young Young = &student
	fmt.Println(young.Name())

	_, ok := interface{}(student).(Young)
	fmt.Printf("Student implements interface Young: %v\n", ok)

	_, ok = interface{}(&student).(Young)
	fmt.Printf("*Student implements interface Young: %v\n", ok)
}

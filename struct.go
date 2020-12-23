package main

import (
	"fmt"
	"reflect"
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

type Named interface {
	Name() string
}

type Young interface {
	SetName(string)
	Named
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

// 值方法的接收者是副本值，在该方法内对该副本的修改一般都不会体现在原值上，除非这个类型本身是某个引用类型（比如切片或字典）的别名类型
// 指针方法的接收者是指针值，在方法内进行修改，会体现在原值上
// 一个数据类型的方法集合中仅包含它的所有值方法，而该类型的指针类型的方法集合包括所有值方法和所有指针方法
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

	// 给一个接口变量赋值的时候，该变量的动态类型会与它的动态值一起被存储在一个专用的数据结构中
	// 这个专用的数据结构叫做 iface 吧，iface 实例包含两个指针，一个是指向类型信息的指针，另一个是指向动态值的指针
	var young Young = &student
	fmt.Println(young.Name())

	// 对于任何数据类型，只要它的方法集合中完全包含了一个接口的全部的方法，那么它就一定是这个接口的实现类型
	_, ok := interface{}(student).(Young)
	fmt.Printf("Student implements interface Young: %v\n", ok)

	_, ok = interface{}(&student).(Young)
	fmt.Printf("*Student implements interface Young: %v\n", ok)

	// 把一个有类型的 nil 赋给接口变量，这个变量的值不会是真正的 nil
	var stu1 *Student
    stu2 := stu1
	var you Young = stu1

	fmt.Printf("stu1 == nil, %v\n", (stu1 == nil))
	fmt.Printf("stu2 == nil, %v\n", (stu2 == nil))
	fmt.Printf("you == nil, %v\n", (you == nil))

	fmt.Printf("you type: %s\n", reflect.TypeOf(you).String())
}

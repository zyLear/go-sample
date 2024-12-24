package main

import "fmt"

type Person struct {
	name string
}

func main() {

	person := Person{
		name: "name1",
	}

	p := &person

	*p = Person{"name2"} // 结构体指针取值只能赋值结构体，且是在原地址赋值，会直接覆盖原结构体

	fmt.Println(person) //{name2}

}

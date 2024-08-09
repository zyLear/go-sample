package main

import (
	"fmt"
)

type Data interface {
	name()
	age()
}

type B interface {
	name()
}

type data1 struct {
}

type data2 struct {
}

func (d *data2) name() {
	fmt.Println("data1 name")
}

func (d *data1) name() {
	fmt.Println("data1 name")
}

func (d *data1) age() {
	fmt.Println("data1 age")

}

func handler(data interface{}) {

	switch value := data.(type) {
	case []interface{}:
		println("inter", value)
		//for _, d := range value {
		//	d.name()
		//}
	//case []*data1:
	//	for _, d := range value {
	//		d.age()
	//	}
	//case []*data2:
	//
	//	for _, d := range value {
	//		d.name()
	//	}
	default:

	}
}

func main() {

	list := []*data1{{}}
	handler(list)

}

package main

import (
	"reflect"
)

type MyInterface interface {
	MyMethod()
}

type MyStruct struct{}

// MyStruct 的指针类型实现了 MyInterface
func (ms *MyStruct) MyMethod() {}

func main() {
	//var myStructPtr MyInterface
	myStructPtr := &MyStruct{}

	// 获取类型
	//typeOfMyStruct := reflect.TypeOf(myStruct)
	//typeOfMyStructPtr := reflect.ValueOf(myStructPtr)
	//typeOfMyStructPtr := reflect.TypeOf((*MyInterface)(nil))
	//var ee MyInterface = &MyStruct{}
	//typeOfMyInterface := reflect.TypeOf(ee)

	dd := &myStructPtr
	types := reflect.TypeOf(dd).Elem()
	//types := reflect.TypeOf(myStructPtr)

	println(types.Kind() == reflect.Ptr)
	println(types.Kind() == reflect.Interface)
	println(types.Kind() == reflect.Struct)
	println(types.Kind())

	//println(typeOfMyInterface.IsNil())

	//// 非指针类型的实现检查
	//fmt.Println(typeOfMyStruct.Implements(typeOfMyInterface))   // false
	//fmt.Println(typeOfMyStruct.AssignableTo(typeOfMyInterface)) // false
	//
	//// 指针类型实现检查
	//fmt.Println(typeOfMyStructPtr.Implements(typeOfMyInterface))   // true
	//fmt.Println(typeOfMyStructPtr.AssignableTo(typeOfMyInterface)) // true
	//
	//// 如果尝试将一个非指针类型的值赋给接口
	//if typeOfMyStruct.Implements(typeOfMyInterface) {
	//	fmt.Println("MyStruct implements MyInterface (with value), but cannot assign")
	//}

}

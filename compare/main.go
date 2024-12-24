package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	a := "{\"a\":1}"
	b := "{\"a\":1,\"b\":2}"
	var objA, objB map[string]interface{}
	err := json.Unmarshal([]byte(a), &objA)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(b), &objB)
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.DeepEqual(a, b))

}

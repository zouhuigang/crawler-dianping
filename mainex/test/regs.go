package main

import (
	"../lib/change"
	"fmt"
	"reflect"
)

type NotknownType struct {
	s1, s2, s3 string
}

var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

type Sy struct {
	Ss1 string
	Ss2 string
	Ss3 string
}

func main() {
	a := Sy{"aaa", "rrr", "sss"}
	map_ := change.Struct.ToMap(a)
	for k, v := range map_ {
		fmt.Println(k, v)
	}
}
func main123() {
	a := Sy{"aaa", "rrr", "sss"}
	map_ := sss(a)
	for k, v := range map_ {
		fmt.Println(k, v)
	}

}

func sss(obj interface{}) map[string]interface{} {
	elem := reflect.ValueOf(obj)
	map_ := map[string]interface{}{}
	for i := 0; i < elem.NumField(); i++ {
		//fmt.Println(elem.Type().Field(i).Name, elem.Field(i).Interface())
		map_[elem.Type().Field(i).Name] = elem.Field(i).Interface()
	}
	return map_
}

func main11() {
	value := reflect.ValueOf(secret)
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}
}

//----------------
type Body struct {
	Person1 string
	Age     int
	Salary  float32
}

func main22() {
	a := Body{"aaa", 2, 12.34}
	elem := reflect.ValueOf(&a).Elem()
	type_ := elem.Type()
	map_ := map[string]interface{}{}
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Name] = elem.Field(i).Interface()
	}
	fmt.Println(map_)
}

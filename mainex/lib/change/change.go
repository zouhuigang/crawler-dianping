package change

import (
	//"fmt"
	//"errors"
	"reflect"
)

type CStruct struct{}

var Struct = CStruct{}
var map_ = map[string]interface{}{}

//结构体转map数组
//a := Sy{"aaa", "rrr", "sss"}
//	map_ := StructToMap(a)
//	for k, v := range map_ {
//		fmt.Println(k, v)
//	}
func (this CStruct) ToMap(obj interface{}) map[string]interface{} {
	elem := reflect.ValueOf(obj)
	type_ := elem.Type()
	//elem := reflect.ValueOf(&obj).Elem()
	//map_ := map[string]interface{}{}
	for i := 0; i < type_.NumField(); i++ {
		//fmt.Println(elem.Type().Field(i).Name, elem.Field(i).Interface())
		map_[type_.Field(i).Name] = elem.Field(i).Interface()
	}
	return map_
}

// 结构体转map，用&形式
func (this *CStruct) ToMapAddr(v interface{}) map[string]interface{} {
	typ := reflect.TypeOf(v)
	val := reflect.ValueOf(v)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	} else {
		return nil
	}

	return this.ToMapAddrValue(val)
}

func (this *CStruct) ToMapAddrValue(val reflect.Value) map[string]interface{} {
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := val.Field(i)
		tpField := typ.Field(i)
		map_[tpField.Name] = field.Interface()
	}
	return map_
}

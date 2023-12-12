package reflection

import (
	"fmt"
	"reflect"
)

func reflectType(a any) {
	t := reflect.TypeOf(a)
	// fmt.Println("Type=", t)
	fmt.Println("Type.Name=", t.Name())
	fmt.Println("Type.Kind=", t.Kind())
}

func reflectValue(a any) {
	v := reflect.ValueOf(a)
	fmt.Println("Value=", v)
	fmt.Println("Value.Kind=", v.Kind())
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		ret := v.Float()
		fmt.Println("ret=", ret)
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		ret := v.Int()
		fmt.Println("ret=", ret)
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		ret := v.Uint()
		fmt.Println("ret=", ret)
	}
}

func reflectSetValue(a any) {
	v := reflect.ValueOf(a)
	fmt.Println("Value=", v)
	fmt.Println("Value.Kind=", v.Kind())
	fmt.Println("Value.Elem=", v.Elem())
	fmt.Println("Value.Elem.Kind=", v.Elem().Kind())
	switch v.Elem().Kind() {
	case reflect.Float32, reflect.Float64:
		v.Elem().SetFloat(2.34)
		ret := v.Elem().Float()
		fmt.Println("ret=", ret)
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		v.Elem().SetInt(100)
		ret := v.Elem().Int()
		fmt.Println("ret=", ret)
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.Elem().SetUint(100)
		ret := v.Elem().Uint()
		fmt.Println("ret=", ret)
	}
}

type Cat struct {
	Name string
	Age  int
}

func (c Cat) Show() {
	fmt.Println("小猫：", c.Name, c.Age)
}

func printField(a any) {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	fmt.Println("NumField=", t.NumField())
	for i := 0; i < t.NumField(); i++ {
		fmt.Println("field type:", t.Field(i))
		fmt.Println("field value:", v.Field(i))
	}
}

func printMethod(a any) {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	fmt.Println("NumMethod=", t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println("method type:", t.Method(i))
		fmt.Println("method value:", v.Method(i))
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func Demo() {
	f := 1.23
	reflectType(f)
	reflectValue(f)
	reflectSetValue(&f)
	n := 10
	reflectType(n)
	reflectValue(n)
	reflectSetValue(&n)
	c := Cat{"c", 0}
	reflectType(c)
	reflectValue(c)
	printField(c)
	printMethod(c)
}

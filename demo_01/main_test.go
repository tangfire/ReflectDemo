package main

import (
	"fmt"
	"reflect"
	"testing"
)

type MyInt int
type User struct {
	IdOrName interface{} `json:"idOrName"`
	M        MyInt       `json:"m"`
}

func TestReflect01(t *testing.T) {
	u := User{
		IdOrName: "tangfire",
		M:        1,
	}

	elem := reflect.ValueOf(&u).Elem()
	elem1 := reflect.ValueOf(u)
	fmt.Println(elem)  // {tangfire 1}
	fmt.Println(elem1) // {tangfire 1}
	name := elem.FieldByName("IdOrName")
	m := elem.FieldByName("M")
	fmt.Printf("type=%s,kind=%s\n", m.Type(), m.Kind())                     // type=main.MyInt,kind=int
	fmt.Printf("type=%s,kind=%s\n", name.Type(), name.Kind())               // type=interface {},kind=interface
	fmt.Printf("type=%s,kind=%s\n", name.Elem().Type(), name.Elem().Kind()) // type=string,kind=string
	name.Set(reflect.ValueOf("myl"))
	fmt.Println(elem) // {myl 1}
	name1 := elem1.FieldByName("IdOrName")
	fmt.Println(name1)
	//name1.Set(reflect.ValueOf("zh")) // Go的反射修改必须通过指针链获取到原始变量的内存地址。通过reflect.ValueOf(&x).Elem()的组合，可以正确获取到可寻址的反射值对象
	fmt.Println(elem1)
}

func TestReflect02(t *testing.T) {
	u := User{
		IdOrName: "tangfire",
		M:        1,
	}

	typeOf := reflect.TypeOf(u)
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		fmt.Println("name:", field.Name)
		fmt.Println("type:", field.Type)
		fmt.Println("kind:", field.Type.Kind())
		fmt.Println("tag:", field.Tag)
		fmt.Println("---------------------")

		//name: IdOrName
		//type: interface {}
		//kind: interface
		//tag: json:"idOrName"
		//---------------------
		//name: M
		//type: main.MyInt
		//kind: int
		//tag: json:"m"
		//---------------------

	}
}

func TestReflect03(t *testing.T) {
	var x float64 = 3.4
	x_v := reflect.ValueOf(x)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", x_v)
	fmt.Printf("%T\n", x_v.Interface())
	fmt.Printf("%T\n", x_v.Interface().(float64))

	//float64
	//reflect.Value
	//float64
	//float64
}

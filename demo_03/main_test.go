package demo_03

import (
	"fmt"
	"reflect"
	"testing"
)

type People interface {
	Color() string
}

type User struct {
	UserName string `json:"userName"`
	age      int    `json:"age"`
	Gender   string `json:"gender"`
}

func (u *User) GetName() string {
	return u.UserName
}

func (u *User) GetAge() int {
	return u.age
}

func (u *User) Add(a, b int) int {
	return a + b
}

func (u *User) Color() string {
	return "black"
}

type Dog struct {
}

func TestReflect(t *testing.T) {
	u := User{UserName: "john"}
	valueOfUser := reflect.ValueOf(u)
	// 获取Name的值
	fmt.Println("Name is ", valueOfUser.FieldByName("UserName"))
}

func TestReflect2(t *testing.T) {
	u := User{UserName: "tangfire"}
	elem := reflect.ValueOf(&u)
	elemv := elem.Elem()
	name := elemv.FieldByName("UserName")
	name.SetString("myl")
	fmt.Println("name is ", name)
	fmt.Println(elem.Type())
	fmt.Println(elemv.Type())

	elema := elemv.Addr()
	fmt.Println(elema.Type())
	fmt.Printf("%T\n", elemv)
	// 从反射对象reflect.Value转为User
	fmt.Printf("%T\n", elemv.Interface().(User))

}

func TestReflect3(t *testing.T) {
	var i interface{}
	v := reflect.ValueOf(i)
	fmt.Println(v.Kind())
	fmt.Println(v.Kind() == reflect.Invalid)

	fmt.Println(v.IsValid())

	var user *User = nil
	u := reflect.ValueOf(user)
	if u.IsValid() {
		fmt.Printf("u持有的值是nil %t\n", u.IsNil()) // 调用IsNil()前先确保IsValid(),否则会panic
	}

	var user1 User
	n := reflect.ValueOf(user1)
	if n.IsValid() {
		fmt.Printf("n持有的值是对应类型的0值 %t\n", n.IsZero()) // 调用IsZero()前先确保IsValid(),否则会panic
	}

}

func TestReflect4(t *testing.T) {
	var user = User{
		UserName: "john",
	}
	valueof := reflect.ValueOf(&user)
	valueof.Elem().FieldByName("UserName").SetString("tangfire")
	fmt.Println(user)
	age := valueof.Elem().FieldByName("age")
	if age.CanSet() {
		age.SetInt(18)
	} else {
		fmt.Println("私有成员不能修改值")
	}

}

func TestReflect5(t *testing.T) {
	var user = User{
		UserName: "john",
	}
	valueof := reflect.ValueOf(&user)
	methodByName := valueof.MethodByName("GetName")
	res := methodByName.Call([]reflect.Value{})
	for _, v := range res {
		fmt.Println(v.Interface().(string))
	}

	res1 := valueof.MethodByName("Add").Call([]reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)})
	for _, v := range res1 {
		fmt.Println(v.Interface().(int))
	}

}

func TestReflect6(t *testing.T) {
	// 创建结构体
	userType := reflect.TypeOf(User{})
	userValue := reflect.New(userType)
	userValue.Elem().FieldByName("UserName").SetString("tangfire")
	fmt.Printf("%T\n", userValue.Elem())
	fmt.Printf("%T\n", userValue.Elem().Interface().(User))
	fmt.Println(userValue.Elem().Interface().(User))

	// 创建切片
	sliceType := reflect.TypeOf([]User{})
	sliceValue := reflect.MakeSlice(sliceType, 1, 3)
	sliceValue.Index(0).FieldByName("UserName").SetString("tangfire")
	fmt.Printf("%T\n", sliceValue.Interface().([]User))
	fmt.Println(sliceValue)
	fmt.Println(sliceValue.Interface().([]User))

}

func TestReflect7(t *testing.T) {
	users := make([]*User, 1, 3)
	users[0] = &User{UserName: "john"}
	usersValue := reflect.ValueOf(&users)
	usersValue.Elem().Index(0).Set(reflect.ValueOf(&User{UserName: "tangfire"}))
	fmt.Println(usersValue.Elem().Index(0))
	usersValue.Elem().SetLen(2)
	usersValue.Elem().Index(1).Set(reflect.ValueOf(&User{UserName: "john216"}))
	fmt.Println(usersValue.Elem().Index(1))
	usersValue = reflect.Append(usersValue.Elem(), reflect.ValueOf(&User{UserName: "fireShine"}))
	u := usersValue.Interface().([]*User)
	fmt.Println(u[2].UserName)
}

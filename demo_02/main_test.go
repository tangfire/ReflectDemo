package demo_02

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
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
}

func (u User) GetName() string {
	return u.UserName
}

func (u *User) GetAge() int {
	return u.Age
}

func (u *User) Add(a, b int) int {
	return a + b
}

func (u *User) Color() string {
	return "black"
}

type Dog struct {
}

func Sub(a, b int) string {
	return fmt.Sprintf("a-b =%d", a-b)
}

func TestReflect(t *testing.T) {
	var user = User{
		UserName: "tangfire",
		Age:      18,
		Gender:   "male",
	}

	typeUser := reflect.TypeOf(user)
	numField := typeUser.NumField()
	for i := 0; i < numField; i++ {
		field := typeUser.Field(i)
		fmt.Printf(
			"%d : name(变量名称)=%s offset(首地址偏移量)=%d \n"+
				"anonymous(是否为匿名变量)=%t type(变量类型)=%s exported(是否可见)=%t tag=%s\n",
			i,
			field.Name,
			field.Offset,
			field.Anonymous,
			field.Type,
			field.IsExported(), // 正确调用方法
			field.Tag,          // 提取json标签
		)
	}

}

// 除了上述方式，也可以使用变量名称获取字段
func TestReflect2(t *testing.T) {

	var user = User{
		UserName: "tangfire",
		Age:      18,
		Gender:   "male",
	}

	typeUser := reflect.TypeOf(user)
	if un, ok := typeUser.FieldByName("UserName"); ok {
		fmt.Printf(
			"name(变量名称)=%s offset(首地址偏移量)=%d \n"+
				"anonymous(是否为匿名变量)=%t type(变量类型)=%s exported(是否可见)=%t tag=%s\n",
			un.Name,
			un.Offset,
			un.Anonymous,
			un.Type,
			un.IsExported(), // 正确调用方法
			un.Tag,          // 提取json标签
		)
	}

}

// 还可以根据索引获取
func TestReflect3(t *testing.T) {
	var user = User{
		UserName: "tangfire",
		Age:      18,
		Gender:   "male",
	}

	typeUser := reflect.TypeOf(user)

	age := typeUser.FieldByIndex([]int{1})
	fmt.Printf(
		"name(变量名称)=%s offset(首地址偏移量)=%d \n"+
			"anonymous(是否为匿名变量)=%t type(变量类型)=%s exported(是否可见)=%t tag=%s\n",
		age.Name,
		age.Offset,
		age.Anonymous,
		age.Type,
		age.IsExported(), // 正确调用方法
		age.Tag,          // 提取json标签
	)
}

func TestReflect4(t *testing.T) {
	var user = User{
		UserName: "tangfire",
		Age:      18,
		Gender:   "male",
	}

	// 不包含指针的方法
	typeOf := reflect.TypeOf(user)
	methodNum := typeOf.NumMethod()
	for i := 0; i < methodNum; i++ {
		method := typeOf.Method(i)
		fmt.Printf("method name:%s ,type:%s ,exported:%t\n", method.Name, method.Type, method.IsExported())
	}

	fmt.Println("----------------------------------")

	// 包含指针或者值的方法
	typeUserPoint := reflect.TypeOf(&user)
	methodNumPoint := typeUserPoint.NumMethod()
	for i := 0; i < methodNumPoint; i++ {
		method := typeUserPoint.Method(i)
		fmt.Printf("method name:%s ,type:%s ,exported:%t\n", method.Name, method.Type, method.IsExported())
	}

}

func TestReflect5(t *testing.T) {
	var user = &User{
		UserName: "tangfire",
		Age:      18,
		Gender:   "male",
	}

	typeOf := reflect.TypeOf(user.Add)
	for i := 0; i < typeOf.NumIn(); i++ {
		fmt.Printf("入参 类型:%s\n", typeOf.In(i).Kind())
	}

	for i := 0; i < typeOf.NumOut(); i++ {
		fmt.Printf("返回值 类型:%s\n", typeOf.Out(i).Kind())
	}

	fmt.Println("--------------------------------------")

	subOf := reflect.TypeOf(Sub)
	for i := 0; i < subOf.NumIn(); i++ {
		fmt.Printf("入参 类型:%s\n", subOf.In(i).Kind())
	}

	for i := 0; i < typeOf.NumOut(); i++ {
		fmt.Printf("返回值 类型:%s\n", subOf.Out(i).Kind())
	}

}

// 判断一个结构体是否实现接口
func TestReflect6(t *testing.T) {
	// 首先我们需要获取到接口的类型 原理是把nil强制转换为*People
	peopleType := reflect.TypeOf((*People)(nil)).Elem()
	fmt.Println("People是否是一个接口:", peopleType.Kind() == reflect.Interface)
	// 判断User和Dog 是否实现了People
	noPointUser := reflect.TypeOf(User{})
	pointUser := reflect.TypeOf(&User{})
	noPointDog := reflect.TypeOf(Dog{})
	pointDog := reflect.TypeOf(&Dog{})
	fmt.Println("noPointUser是否实现了接口:", noPointUser.Implements(peopleType))
	fmt.Println("pointUser是否实现了接口:", pointUser.Implements(peopleType))
	fmt.Println("noPointDog是否实现了接口:", noPointDog.Implements(peopleType))
	fmt.Println("pointDog是否实现了接口:", pointDog.Implements(peopleType))

	//People是否是一个接口: true
	//noPointUser是否实现了接口: false
	//pointUser是否实现了接口: true
	//noPointDog是否实现了接口: false
	//pointDog是否实现了接口: false
}

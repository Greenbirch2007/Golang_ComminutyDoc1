package main

import (
	"fmt"
	"reflect"
)

//student结构体

type Student struct {
	Name string `json:"name1" form:"sssdd"`
	Age int `json:"age"`
	Score int `json:"score"`
}

func (s Student) GetInfo() string{
	var str = fmt.Sprintf("%v  %v  %v",s.Name,s.Age,s.Score)
	return str
}

func (s *Student)SetInfo(name string,age int,score int){
	s.Name =name
	s.Age=age
	s.Score = score
}

func (s Student) Print(){
	fmt.Println("这是一个打印方法")
}



//打印字段

func PrintStructField(s interface{}){
	//1. 通过类变量里面的Field可以获取结构体的字段
	// 2. 通过类型变量里面的FieldByName可以获取结构体的字段
	// 3. 通过类型变量里面的NumField获取到该结构体有几个字段

	//判断参数是不是结构体类型
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct{
		fmt.Println("传入的参数不是一个结构体")
		return
	}
	fileld0:=t.Field(0)//小米
	fmt.Printf("%#v \n",fileld0)//reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0x4adf20), Tag:"json:\"name\"", Offset:0x0, Index:[]int{0}, Anonymous:false}
	fmt.Println("字段名称",fileld0.Name)
	fmt.Println("字段类型",fileld0.Type)
	fmt.Println("字段Tag",fileld0.Tag.Get("json"))
	//如果想要获取一个未知的结构体的字段名，就得用到反射
	fmt.Println("字段Tag",fileld0.Tag.Get("form"))

	// 2. 通过类型变量里面的FieldByName可以获取结构体的字段
	fmt.Println("---------------------------")
	field1,ok := t.FieldByName("Age")
	if ok{
		fmt.Printf("%#v \n",field1)//reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0x4adf20), Tag:"json:\"name\"", Offset:0x0, Index:[]int{0}, Anonymous:false}
		fmt.Println("字段名称",field1.Name)
		fmt.Println("字段类型",field1.Type)
		fmt.Println("字段Tag",field1.Tag.Get("json"))
	}

	// 3. 通过类型变量里面的NumField获取到该结构体有几个字段

	var fieldCount = t.NumField()
	fmt.Println("结构体有",fieldCount,"属性")

	//4.通过值变量获取结构体属性对应的值
	fmt.Println(v.FieldByName("Name"))
	fmt.Println(v.FieldByName("Age"))
	//使用for循环来一次性返回结构体的属性名称，属性值，属性类型，属性Tag
	fmt.Println("---------------------------")
	for i:=0;i<fieldCount ;i++  {
		fmt.Printf("属性名称:%v  属性值:%v   属性类型%v  属性Tag %v",t.Field(i).Name,v.Field(i),t.Field(i).Type,t.Field(i).Tag.Get("json"))

	}



}

//打印执行方法

func PrintStructFn(s interface{}){
	// 1.  通过类型变量里面的Method可以获取结构体的方法
	// 2. 通过类型变量获取这个结构体有多少个方法
	// 3. 通过(值变量)执行方法(注意需要使用值变量，并且要注意参数)
	// v.Method(0).Call(nil)

	//4. 执行方法传入参数(注意需要使用值变量)，并且需要参数，
	//接收的参数是[]reflect.Value
	//5. 执行方法获取方法的值


	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct{
		fmt.Println("传入的参数不是一个结构体")
		return
	}

	// 1.  通过类型变量里面的Method可以获取结构体的方法

	method0 := t.Method(0) // 和结构体方法的顺序没有关系,和结构体方法的ASCII有关
	fmt.Println(method0.Name)
	fmt.Println(method0.Type)
	// 2. 通过类型变量获取这个结构体有多少个方法
	fmt.Println("---------------------------")

	method1,ok   := t.MethodByName("Print")
	if ok{
		fmt.Println(method1.Name)
		fmt.Println(method1.Type)
	}
	fmt.Println("---------------------------")

	// 3. 通过(值变量)执行方法(注意需要使用值变量，并且要注意参数)
	//v.Method(0).Call(nil)
	//v.Method(1).Call(nil)
	//nil表示不传入任何参数
	v.MethodByName("Print").Call(nil)
	info1 :=v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info1)

	//4. 执行方法传入参数(注意需要使用值变量)，并且需要参数，
	//接收的参数是[]reflect.Value
	var params []reflect.Value //这里通过反射，改变了结构体属性的值
	params = append(params,reflect.ValueOf("李四"))
	params = append(params,reflect.ValueOf(26))
	params = append(params,reflect.ValueOf(98))
	v.MethodByName("SetInfo").Call(params) //执行方法传入参数
	info2 :=v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info2)

	//5. 获取方法的数量
	fmt.Println(t.NumMethod())


}

func reflectChangeStruct(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Ptr { //这里判断是不是指针类型
		fmt.Println("传入的不是结构体指针类型")
		return
	} else if t.Elem().Kind() != reflect.Struct {
		//这里判断是不是结构体指针类型
		fmt.Println("传入的不是结构体指针类型")
		return

	}
	//修改结构体属性的值
	name := v.Elem().FieldByName("Name")
	name.SetString("小李")
	age := v.Elem().FieldByName("Age")
	age.SetInt(20)
}






func main(){

	stu1 := Student{
		Name:  "小米",
		Age:   16,
		Score: 99,
	}

	//PrintStructField(stu1)
	//PrintStructFn(&stu1)

	reflectChangeStruct(&stu1)
	fmt.Println(stu1)



}
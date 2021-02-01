package datatype

import (
	"fmt"
	"time"
)

/**
结构体有点类似于java中的bo类

*/

func TypeStructOut() {

}

/**
1. 定义了一个全局的结构体及其实例化变量
*/
type employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

/**
实例化
*/
func typeInit() {

	//1. 普通实例化
	type User struct {
		Id   int
		Name string
	}

	var a User
	a.Id = 1
	b := User{1, "Linken"}
	fmt.Println(b)

	// 2. 指针实例化
	// 通过new(Type)或者$T{}
	c := new(User)
	fmt.Println(c)
	d := &User{1, "fasdf"}
	fmt.Println(d)

	var linken employee

	// 1.1 通过"."操作符可以修改结构体变量中的
	linken.Salary -= 5000

	// 1.2 通过"&"操作符取地址；通过"*"操作符获取对应地址的变量值
	// "*"操作符作用于地址类变量，表示取对应地址的值；作用于Type上时，表示是该类型为指针类型
	position := &linken.Position
	*position = "Senior" + *position

	emp := &linken
	emp.Position += "asdfasdf"
	(*emp).Position += "adfasdfas"

}

/**
3. 结构体的比较，可以通过==进行比较，会比较结构体的每个成员

可比较的类型可以作为map的key
*/
func compareStruct() {
	// 3.1 结构体可进行比较
	type Point struct {
		X, Y int
	}

	p := Point{1, 2}
	q := Point{2, 1}

	// 如下的比较是等价的
	fmt.Println(p.X == q.X && p.Y == q.Y)
	fmt.Println(p == q)

	//3.2 结构体是可比较类型，因此可以作为map的key
	type Address struct {
		hostname string
		port     int
	}

	hits := make(map[Address]int)
	hits[Address{"baidu.com", 443}]++
}

/***********************************************
4. 结构体嵌入和匿名成员
下面通过一个例子进行展示
***********************************************/

// 4.1 基本类型定义，圆：圆心+半径；轮形：圆心+半径+径向辐条、
// 下面展示了圆形和轮形的结构体定义以及轮形的实例化
func extends() {

	type Circle struct {
		X, Y, Radius int
	}

	type Wheel struct {
		X, Y, Radius, Spokes int
	}

	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
}

// 4.2 基本结构抽象优化，可以将圆心，圆形等抽象成公共的类型

func extendsAgain() {
	type Point struct {
		X, Y int
	}

	type Circle struct {
		Center Point
		Radius int
	}

	type Wheel struct {
		Circle Circle
		Spokes int
	}

	// 改动后结构体虽然清晰了，但是访问起来特别繁琐
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20
}

// 4.3 匿名成员：声明一个成员对应的数据类型 而 不指定成员变量的名字；匿名成员的数据类型必须是命名的类型或者是一个命名的类型的指针
func unknownName() {

	type Point struct {
		X, Y int
	}

	type Circle struct {
		Point
		Radius int
	}

	type Wheel struct {
		Circle
		Spokes int
	}

	// 可以通过以下方式进行访问
	var wh Wheel
	wh.X = 8      // equivalent to w.Circle.Point.X = 8
	wh.Y = 8      // equivalent to w.Circle.Point.Y = 8
	wh.Radius = 5 // equivalent to w.Circle.Radius = 5
	wh.Spokes = 20
}

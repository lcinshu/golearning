package datatype

import "fmt"

// ------------------- 数组 ----------------------

/**
一、数组初始化
*/
func Init() {

	/**
	一、
	数组类型 = 数组大小 + 数组元素类型
	两者缺一不可
	*/
	var a [3]int
	b := [3]int{1, 2, 3}
	b = [...]int{1, 7, 3}

	/**
	二、
	不同大小，不同类型的数组无法赋值
	例如下面的就无法复制给b
	b = [4]int{2, 3, 4, 5}
	*/

	fmt.Printf("%d \n\n", &a)
	fmt.Printf("%d \n\n", b)

	/**
	三、
	顺序定义，将数组按照值顺序进行定义
	指定一个索引和对应值列表的方式初始化iota

	*/
	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbo := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbo[RMB])

	/**
	四、
	批量定义
	如下初始化语句表示：前10个元素初始化为0，最后一个元素初始化为-1；总共11个元素
	*/
	r := [...]int{10: -1}

	fmt.Println(r)

}

/**
二、数组的比较
数组元素如果可以比较，则数组类型也可以比较
可以直接通过==符号进行比较
当且仅当，数组的所有元素都相等的时候数组才是相等的，反之则为不相等
*/
func Compare() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	fmt.Println(a != b, a != c, b != c)
	d := [3]int{1, 2}
	fmt.Printf("%T\n", d)
	// fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
}

/**
print
%x:指定以十六进制的格式打印数组或slice全部的元素
%t:用于打印布尔型数据
%T:用于显示一个值对应的数据类型
*/
func Array() {
	a := []int{1, 2, 3}

	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	fmt.Printf("%x\n", a)
	fmt.Printf("%t\n", false)
	fmt.Printf("%T\n", a)

	// print the index and value of array a
	for i, v := range a {
		fmt.Printf("%d %d \n\n\n", i, v)
	}

	// only print the value of array a
	for _, v := range a {
		fmt.Printf("%d \n\n", v)
	}

	// only print the index of array a
	for i, _ := range a {
		fmt.Printf("%d \n\n", i)
	}

}

func Index(ptr *[32]byte) {
	*ptr = [32]byte{}

	for i := range ptr {
		ptr[i] = 0
	}
}

package complextype

import (
	"fmt"
	"reflect"
)

/**
Slice（切片）代表变长的序列，序列中每个元素都有相同的类型。一个slice类型一般写作[]T，其中T代表slice中元素的类型；slice的语法和数组很像，只是没有固定长度而已

——摘自go语言圣经

关于slice的详解，可参考：https://www.cnblogs.com/f-ck-need-u/p/9854932.html
*/

func SliceToOut() {
	// 1. slice基本性质了解
	//SliceInit()

	// 2. 反转整个数组
	//reverseSlice()

	// 3. copy
	copySlice()
}

/**
slice基本性质
*/
func SliceInit() {
	// 0、 array和slice在定义上的显式区别，是否有定义长度
	// 0.1、数组可以使用 "=" 进行相等判断，而slice不能使用 "=" 进行判断
	monthArray := [13]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June",
		7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	month := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June",
		7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}

	// 1. 打印数组元素，长度和容量
	fmt.Println("monthArray：", monthArray)
	fmt.Println("month元素：", month, month[:], month[1:])
	fmt.Println("month长度和容量", len(month), cap(month))

	// 2. 数组切片
	// 2.1 切片是在原slice基础上新增一个指针而已；在切片上对数组进行修改，会影响原Slice的数据
	Q2 := month[4:7]
	summer := month[1:6]
	summer[0] = "ASDFAS"

	endlessSummer := summer[:11]

	fmt.Println("四到六月：", Q2)
	fmt.Println("夏季：", summer)
	//fmt.Println("数组切片，超过原slice容量上限就会报错：", summer[:20])

	// 2.2 切片都是在底层数组基础上，定义别名，因此地址都是同一个
	fmt.Println(&summer[0])
	fmt.Println(&endlessSummer[0])
	fmt.Println("数组切片，超过当前子slice容量上限，但仍在原slice容量范围内，则会扩容子序列：", endlessSummer)
	fmt.Println("原month数组：", month)
}

/**
1. 翻转slice
2. 从第N个元素开始左右翻转slice
*/
func reverseSlice() {
	// 3、翻转slice
	// 3.1、翻转整个slice
	reverseSlicea := []int{1, 2, 3, 4, 5, 6, 7}
	reverse(reverseSlicea)
	fmt.Println("反转之后的数组：", reverseSlicea)

	// 3.2 以第N个元素为轴进行翻转，比如以第三个元素为轴，想做进行翻转
	reverseSliceb := []int{1, 2, 3, 4, 5, 6, 7}
	reverse(reverseSliceb[:3])
	reverse(reverseSliceb[3:])
	reverse(reverseSliceb)
	fmt.Println(reverseSliceb)

	returnBackd := []int{0, 2, 1, 2, 0}
	returnBackNew := returnBack(returnBackd)
	fmt.Println(reflect.DeepEqual(returnBackd, returnBackNew))
}

/**
这表示将src slice拷贝到dst slice，src比dst长，就截断，src比dst短，则只拷贝src那部分。
*/
func copySlice() {
	s1 := []int{11, 22, 33}
	s2 := make([]int, 5)
	s3 := make([]int, 2)

	num := copy(s2, s1) // 返回被拷贝的元素数
	copy(s3, s1)

	fmt.Println(num) // 3
	fmt.Println(s2)  // [11,22,33,0,0]
	fmt.Println(s3)  // [11,22]
}

/**
比较两个数组是否相同
*/
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

/**
反转数组
设置两个指针i,j，一头一尾，一个递增，一个递减，对应位置进行交换
*/
func reverse(sli []int) {
	for i, j := 0, len(sli)-1; i < j; i, j = i+1, j-1 {
		sli[i], sli[j] = sli[j], sli[i]
	}
}

/**
不太合适的回文子串
*/
func returnBack(sli []int) []int {
	for i, j := 0, len(sli)-1; i < j; i, j = i+1, j-1 {
		sli[i], sli[j] = sli[i], sli[i]
	}
	return sli
}

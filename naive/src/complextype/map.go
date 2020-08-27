package complextype

import (
	"fmt"
	"sort"
)

func MapOut() {

}

/**
map 基本结构定义
*/
func mapInit() {

	// 1. 定义数组
	ages1 := make(map[string]int)
	ages1["alice"] = 31
	ages1["jack"] = 33
	fmt.Println(ages1)
	fmt.Println(ages1["alice"])

	ages2 := map[string]int{
		"jack":       23,
		"alice":      22,
		"yanhong li": 43,
		"jack ma":    44,
		"tony ma":    45,
		"xfgqwef":    32,
	}
	fmt.Println(ages2)

	// 2. 增删改查
	delete(ages2, "jack ma")
	ages2["bob"] = ages2["bob"] + 1
	ages2["bob"] += 1
	ages2["bob"]++

	// 3. 遍历：map的遍历是随机的，每次遍历的结果都不一样
	for name, age := range ages2 {
		fmt.Printf("%s\t%d\n\n", name, age)
	}

	// 4. 排序
	// 4.1 因为知道ages的大小，所以我们可以通过make直接指定新数组的大小
	// 4.2 第一个排序中，只需要key的值，因此值收集第一个参数，第二个参数可以省略

	names := make([]string, 0, len(ages2))
	for name := range ages2 {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n\n", name, ages2[name])
	}

}

package datatype

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

	// 0. 空数组
	var nilMap map[string]int
	fmt.Println(nilMap == nil)
	fmt.Println(len(nilMap) == 0)
	//nilMap["alice"]=32 // nil数组无法新增数据

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

	// 5. 判断元素是否存在
	// 5.1 如果key对应的元素存在，则返回该key对应的value；如果元素不存在，则返回该value对应类型的零值
	// 5.2 map的下标语法会返回两个值；第二个值标识元素是否存在；false：不存在；true：存在
	if age, ok := ages2["alice"]; ok {
		fmt.Println(age)
	}
}

/**
比较两个map是否相同
map之间不能通过等号进行比较（除了和nil进行比较之外）
*/
func equalMap(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	fmt.Sprintf("%q", x)
	return true
}

/**
map的key需要支持比较，如果map的key是复杂类型一般不支持直接比较，需要转换成基本类型作为key（key仅作为查找之用，不作为后续计算依据）
下面的例子演示了如何使用map来记录提交相同的字符串列表的次数
*/
var m = make(map[string]int)

func k(list []string) string { return fmt.Sprintf("%q", list) }

func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }

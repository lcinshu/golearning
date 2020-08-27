package complextype

import "time"

/**
结构体有点类似于java中的bo类

*/

func TypeStructOut() {

}

/**
定义了一个全局的结构体变量
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

var linken employee

/**
结构体基本属性
*/
func typeInit() {
	linken.Salary -= 5000

	position := &linken.Position
	*position = "Senior" + *position
}

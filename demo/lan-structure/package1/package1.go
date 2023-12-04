package p1
import "fmt"

// 首字母大写的变量为公有变量 可以被外部引用 
// 首字母小写的变量为私有变量 只能在当前包内使用
// 公有函数 首字母大写
func F1() {
	fmt.Println("调用package1的 F1函数")
}
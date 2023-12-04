// 声明包
package main

// 引入包
import (
	"fmt"
	"lan-structure/package1"
	_ "lan-structure/package2"
)

// 初始化函数，golang每个包的引用会优先调用该函数
func init() {
	fmt.Println("init function")
}


// main函数为整个程序的入口
func main() {
	
	// 语句和表达式
	fmt.Println("Hello World!")
	// 变量
	var a int = 10
	fmt.Printf("打印参数：%d\n", a)
	// for循环
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// 包名可以与包引用目录不一致
	p1.F1()
}


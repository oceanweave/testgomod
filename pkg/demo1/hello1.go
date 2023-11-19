// 更改包名 保持和目录名一致  这样 import 时候不会产生歧义
package demo1

import "fmt"

// v1 版本
func Hello1() {
	fmt.Println("hello1 go package v1 bug-fix ")
	fmt.Println("this is Hello-1 from demo1")
}

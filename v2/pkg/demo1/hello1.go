// 更改包名 保持和目录名一致  这样 import 时候不会产生歧义
package demo1

import "fmt"

// v2 发生巨大重构
func Hello1() {
	fmt.Println("has big change in this demo1 package, totallly different from v1 ")
	fmt.Println("this is Hello-1 from demo1")
}

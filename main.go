package main

import (
	// 2-1 引用新改的包名 demox，发现 go run main.go 无法运行
	//"github.com/oceanweave/testgomod/pkg/demox"

	// 2-2 引用包的路径，发现可以运行
	// github.com/oceanweave/testgomod 是项目名，go.mod 的名称，可以理解为当前项目的根目录
	// 因此 github.com/oceanweave/testgomod/pkg/demo1 就可以指向 demox 包
	// 之后下面就可以调用 demox 包

	// 2-3 利用包的重命名
	// 由于此处末尾是 demo1，而该路径下的包是 demox
	// 容易造成如下混乱：1. 不好理解，为什么上面是 demo1，而 demox 执行正常 2. 不便于排错，若 demox 的函数出现问题，不熟悉代码的人不清楚该包所在位置
	// 所以此处，为该 import 进行重命名，标识，以便容易理解

	// 2-4 最终推荐，还是保持包名和目录名一致，这样没有歧义，便于理解和排错
	"github.com/oceanweave/testgomod/pkg/demo1"
)

func main() {
	// 上面 import 引用正确的路径后，就可以正常调用执行
	// 注意更改包名后，可能此处调用会报错，因此项目还没有刷新过来，此时关闭并重新打开该项目，就可以变为正常了
	demo1.Hello1()
}

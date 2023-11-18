package main

import (
	// 2-1 引用新改的包名 demox，发现 go run main.go 无法运行
	//"github.com/oceanweave/testgomod/pkg/demox"

	// 2-2 引用包的路径，发现可以运行
	// github.com/oceanweave/testgomod 是项目名，go.mod 的名称，可以理解为当前项目的根目录
	// 因此 github.com/oceanweave/testgomod/pkg/demo1 就可以指向 demox 包
	// 之后下面就可以调用 demox 包
	"github.com/oceanweave/testgomod/pkg/demo1"
)

func main() {
	// 上面 import 引用正确的路径后，就可以正常调用执行
	demox.Hello1()
}

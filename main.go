package main

import (
	// 引用新改的包名 demox
	"github.com/oceanweave/testgomod/pkg/demox"
)

// 引用新改的包名，发现 go run main.go 无法运行
func main() {
	demox.Hello1()
}

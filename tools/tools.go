package tools

import "fmt"

// 此 commit 提交后，准备发布版本，如何发布呢？
//  1. 提交 commit，获取 commit-id
//  2. 本地打 tag， git tag -a v1.0.1 commit-id -m "Release v1.0.1"
//     需要注意 go 的版本规则
//     版本形式要求 https://go.dev/ref/mod#versions
//     - 形式建议 v1.2.3, 其他形式可能会出错，导致无法拉取
//     - 1 主版本号：发布了不兼容的版本迭代时递增（breaking changes）。
//     - 2 次版本号：发布了功能性更新时递增。
//     - 3 修订号：发布了bug修复类更新时递增。
//  3. 发布到 github 上 git push origin v1.0.1
func SayHello() {
	fmt.Println("this sayHello from testgomod Repo‘s tools go package")
}

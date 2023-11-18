// 假设目前所有架构都进行了重构，完全不同于 v1 版本，以往使用 v1 的客户使用目前 v2 版本将会出错
// 所以此处更改 module 名称，避免用户应用错误

  //因为 v2 进行了重大重构，v1 与 v2 不兼容，所以更改此处 module 进行重命名，避免用户拉取后的不兼容导致失败
  //若 v2 和 v1 兼容，便可不改此处名称，之后发布 v2.0.0 版本，v2 也可以通过 如下方式拉取
	//兼容版本发布的拉取形式 go get -u github.com/oceanweave/testgomod@v2.0.0
  //
  //不兼容，更改module名称后，现在拉取方式为
	//v1 版本拉取 go get -u github.com/oceanweave/testgomod@v1.0.1
	//v2 版本拉取 go get -u github.com/oceanweave/testgomod/v2@v2.0.0
  //import 使用时
	//v1 版本 import "github.com/oceanweave/testgomod/tools"
  //v2 版本 import "github.com/oceanweave/testgomod/v2/tools"
 	//若同时使用，可利用重命名机制
  //v1 版本重命名为 import v1tools "github.com/oceanweave/testgomod/tools"
module github.com/oceanweave/testgomod/v2

go 1.21.4

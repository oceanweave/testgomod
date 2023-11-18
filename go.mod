// 假设目前所有架构都进行了重构，完全不同于 v1 版本，以往使用 v1 的客户使用目前 v2 版本将会出错
// 所以此处更改 module 名称，避免用户应用错误
module github.com/oceanweave/testgomod/v2

go 1.21.4

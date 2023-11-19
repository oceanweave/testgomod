# 总结

- 开发注意
  - import 引用的是包的路径，而不是包的名称
  - go mod init module-name（ module-name 最好设置为 github 上 repo 路径，这样便于后续引用）
    - 如 module-name 为 `github.com/{your-username}/{your-repo-name}`
    - 此项目就是  `github.com/oceanweave/testgomod`
  - 版本不兼容的时候，可以更改 module-name，以用于区别已发布的版本
    - 如现在 v2 与 v1 不兼容，更改 module-name 为 `github.com/oceanweave/testgomod/v2`
    - 同样 go get 和 import 方式也进行了改动，增加了 v2 路径，下面有详细说明
  - 项目内每个子目录可以视为一个包，建议包名和子目录名保持一致，这样引用时没有歧义，也便于定位错误
    - 当前项目中 main 函数，引用子目录的包时，可以采用 `import moudule-name/包的与该项目根目录的相对位置`
    - 如 demo1 包，`github.com/oceanweave/testgomod/pkg/demo1`, 因此可以理解为 module-name 就是项目的根路径
    - 同样发布到 github 后，别人引用时，也要指定包的【详细路径】
      - 如 `import  github.com/oceanweave/testgomod`, 会发生错误，因为这是一个大项目，有很多 go 包，没有指定详细路径的话，就不知道该如何使用
      - 因此需要指定要使用的包，如使用 demo1 包，`import github.com/oceanweave/testgomod/pkg/demo1`，之后便可以利用 `demo1.函数名` 调用相应的函数
      - 同时注意需要先下载，下载时指定的是整个项目，而不是某个包 `go get -u github.com/oceanweave/testgomod`， -u 参数表示获取最新的 commit 的 go 包

- 因为 v2 进行了重大重构，v1 与 v2 不兼容，所以更改此处 module 进行重命名，避免用户拉取后的不兼容导致失败
  - v1 版本的 module 名称 `module github.com/oceanweave/testgomod`
  - v2 版本的 module 名称 `module github.com/oceanweave/testgomod/v2`

- 若 v2 和 v1 兼容，便可不改此处 module 名称，之后发布 v2.0.0 版本，v2 也可以通过 如下方式拉取
  - 兼容版本发布的拉取形式 go get -u github.com/oceanweave/testgomod@v2.0.0

- 不兼容，更改module名称后，现在拉取方式为
  - v1 版本拉取 go get -u github.com/oceanweave/testgomod@v1.0.1
  - v2 版本拉取 go get -u github.com/oceanweave/testgomod/v2@v2.0.0
- import 使用时
  - v1 版本 import "github.com/oceanweave/testgomod/tools"
  - v2 版本 import "github.com/oceanweave/testgomod/v2/tools"
- 若同时使用，可利用重命名机制
  - v1 版本重命名为 import v1tools "github.com/oceanweave/testgomod/tools"

- 参考自
    - https://www.liwenzhou.com/posts/Go/package/

- 发布版本，第一步，打 tag
    - `git tag -a 版本号 commit-id -m 附加信息`
    - ` git tag -a v1.0 de7cc24 -m "Release v1.0" `
- 发布版本，第二步，发布到 github 上
    - `git push origin v1.0.1`
- 若提交到 main，可以省略分支名称
    - `git push origin v1.0`

- 发布版本，版本形式要求 https://go.dev/ref/mod#versions
    - 形式建议 v1.2.3
    - 1 主版本号：发布了不兼容的版本迭代时递增（breaking changes）。
    - 2 次版本号：发布了功能性更新时递增。
    - 3 修订号：发布了bug修复类更新时递增。

- 发布版本
    - git tag -a v1.0.1 0c74f12d -m "Release v1.0.1"
    - git push origin v1.0.1


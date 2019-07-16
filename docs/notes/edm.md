# Go 外部依赖管理

> External dependency management

## govendor

长期以来，golang 对外部依赖都没有很好的管理方式，只能从 \$GOPATH 下查找依赖。这就造成不同用户在安装同一个项目适合可能从外部获取到不同的依赖库版本，同时当无法联网时，无法编译依赖缺失的项目。

自 1.5 版本开始引入 govendor 工具，该工具将项目依赖的外部包放到项目下的 vendor 目录下（对比 nodejs 的 node_modules 目录），并通过 vendor.json 文件来记录依赖包的版本，方便用户使用相对稳定的依赖。

- 对于 govendor 来说，主要存在三种位置的包：
  - 项目自身的包组织为本地（local）包；
  - 传统的存放在 \$GOPATH 下的依赖包为外部（external）依赖包；
  - 被 govendor 管理的放在 vendor 目录下的依赖包则为 vendor 包。
- 具体来看，这些包可能的类型如下：

| 状态      | 缩写状态 | 含义                                               |
| --------- | -------- | -------------------------------------------------- |
| +local    | l        | 本地包，即项目自身的包组织                         |
| +external | e        | 外部包，即被 \$GOPATH 管理，但不在 vendor 目录下   |
| +vendor   | v        | 已被 govendor 管理，即在 vendor 目录下             |
| +std      | s        | 标准库中的包                                       |
| +unused   | u        | 未使用的包，即包在 vendor 目录下，但项目并没有用到 |
| +missing  | m        | 代码引用了依赖包，但该包并没有找到                 |
| +program  | p        | 主程序包，意味着可以编译为执行文件                 |
| +outside  |          | 外部包和缺失的包                                   |
| +all      |          | 所有的包                                           |

- 常见的命令如下，格式为 govendor COMMAND。通过指定包类型，可以过滤仅对指定包进行操作。

| 命令         | 功能                                                           |
| ------------ | -------------------------------------------------------------- |
| init         | 初始化 vendor 目录                                             |
| list         | 列出所有的依赖包                                               |
| add          | 添加包到 vendor 目录，如 govendor add +external 添加所有外部包 |
| add PKG_PATH | 添加指定的依赖包到 vendor 目录                                 |
| update       | 从 \$GOPATH 更新依赖包到 vendor 目录                           |
| remove       | 从 vendor 管理中删除依赖                                       |
| status       | 列出所有缺失、过期和修改过的包                                 |
| fetch        | 添加或更新包到本地 vendor 目录                                 |
| sync         | 本地存在 vendor.json 时候拉去依赖包，匹配所记录的版本          |
| get          | 类似 go get 目录，拉取依赖包到 vendor 目录                     |

## gomod

从 1.11 版(及其以上版本)开始 go 加入了环境变量`GO111MODULE` 默认 `GO111MODULE=auto`(auto 是指如果在 gopath 下不启用 mod)

- `go mod help` 查看帮助
- `go mod init <项目模块名称>`初始化模块，会在项目根目录下生成 go.mod 文件。
- `go mod tidy` 根据 go.mod 文件来处理依赖关系。
- `go mod vendor` 将依赖包复制到项目下的 vendor 目录。建议一些使用了被墙包的话可以这么处理，方便用户快速使用命令 go build -mod=vendor 编译
- `go list -m all` 显示依赖关系。go list -m -json all 显示详细依赖关系。
- `go mod download <path@version>`下载依赖。参数`<path@version>`是非必写的，path 是包的路径，version 是包的版本。

### 用法

- 在 gopath 外新建一个项目，单独开一个 cmd 设置 set GO111MODULE=on(习惯性的和 git 初始化一样)`go mod init`然后报错了。 正解如下：`go mod init xxx（module名称可与文件名不同）`
- 在项目目录下执行 go mod tidy 下载完成后项目路径下会生成 go.mod 和 go.sum</br>
  go.mod 文件必须要提交到 git 仓库，但 go.sum 文件可以不用提交到 git 仓库(git 忽略文件.gitignore 中设置一下)。
- go 模块版本控制的下载文件及信息会存储到 GOPATH 的 pkg/mod 文件夹里。
- 在国内访问 golang.org/x 的各个包都需要翻墙，我们可以在 go.mod 中使用 replace 替换成 github 上对应的库。</br>
  **强烈建议使用代理：<https://goproxy.io> 添加 go 环境变量 `export GOPROXY=https://goproxy.io`**

### 坑

在引用 mongodb 的包时候报错
go mod labix.org/v2/mgo@v0.0.0-20140701140051-000000000287: bzr branch --use-existing-dir
解决办法 谷歌论坛

在 go.mod 中

```go.mod
replace (
labix.org/v2/mgo => github.com/go-mgo/mgo v0.0.0-20160801194620-b6121c6199b7
launchpad.net/gocheck => github.com/go-check/check v0.0.0-20180628173108-788fd7840127
)
```

- 引入本地包的方法(在 go.mod 中)

```go.mod
require (
test v0.0.0
)

replace (
test => ../test
)
```

> 注意:

1. 引入的包必须也是 gomod 的(有.mod 文件)
2. replace 时必须使用相对路径和 http 路径比如`../` `./` `github.com/go-mgo/mgo`
3. require 的包后必须带版本号
4. go.mod 文件必须传入 git 服务器上

### GOPROXY

- <https://goproxy.io>
- <https://goproxy.cn>

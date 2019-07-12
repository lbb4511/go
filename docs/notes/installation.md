# 环境搭建

## 安装 go

### MAC

```sh
brew install go
```

### linux

- Debian/Ubuntu 系 `sudo apt-get install golang`
  ~~- 卸载 golang-go 及其依赖并删除其本地和配置文件 `sudo apt-get purge --auto-remove golang-go`~~
- Arch 系`sudo pacman -S go`
- Redhat 系`sudo yum install golang`
- 其他

  ```bash
  wget https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz
  tar -C /usr/local -zxf go1.6.2.linux-amd64.tar.gz
  ```

- 环境变量 /etc/profile

  ```profile
  export GOPATH=/home/lbb/work/gopackage
  export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
  ```

  > 注：go 默认环境变量 GOROOT:`GOROOT=/usr/local/go`GOPATH:`GOPATH=~/go` > **_Linux 系统安装到/usr/local/不用设置\$GOROOT,其他路径要设置_**

### Windows

- 到[官网](https://golang.google.cn/)去下需要的安装包

## 开发环境

### VS code

安装 go 插件

### Sublime Text3

> - [Golang 开发环境搭建-Vim 篇](http://studygolang.com/articles/1785)

- [github 的 vim-go-ide](https://github.com/farazdagi/vim-go-ide)

### VIM

> - [Golang 开发环境搭建-Vim 篇](http://studygolang.com/articles/1785)

- [把 vim 当做 golang 的 IDE](http://www.cnblogs.com/zhangqingping/p/5215071.html)
- [基于 vim 搭建 Go 开发环境](http://blog.csdn.net/chosen0ne/article/details/40782991)
- [vim 安装 vim-go 打造 GOLANG 专用 IDE](http://studygolang.com/articles/2927)
- [为 Vim 配置 Golang 开发环境](http://ju.outofmemory.cn/entry/89207)

### Atom

> - [打造 atom 成为 golang 开发神器](http://blog.csdn.net/sweetvvck/article/details/50333327)

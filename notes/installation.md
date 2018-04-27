# 环境搭建

- 安装 `sudo apt-get install golang`
- 到[官网](https://golang.org/dl/)去下需要的安装包
```golang
wget https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz
tar -C /usr/local -zxf go1.6.2.linux-amd64.tar.gz
```

### [下载](http://www.golangtc.com/download) 本地路径(/usr/local/go)放在这个路径下就不用设GOROOT变量了。
    - 卸载golang-go `sudo apt-get  remove golang-go`
    - 卸载golang-go及其依赖 `sudo apt-get remove --auto-remove golang-go`
    - 卸载golang-go并删除其本地和配置文件 `sudo apt-get purge golang-go`
    - 卸载golang-go及其依赖并删除其本地和配置文件 `sudo apt-get purge --auto-remove golang-go`

### 环境变量 /etc/profile
```
export GOPATH=/home/lbb/work/gopackage
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```
***Linux系统安装到/usr/local/不用设置$GOROOT,其他路径要设置***

##### Sublime Text3
>- [Golang开发环境搭建-Vim篇](http://studygolang.com/articles/1785)
- [github的vim-go-ide](https://github.com/farazdagi/vim-go-ide)

##### VIM
>- [Golang开发环境搭建-Vim篇](http://studygolang.com/articles/1785)
- [把vim当做golang的IDE](http://www.cnblogs.com/zhangqingping/p/5215071.html)
- [基于vim搭建Go开发环境](http://blog.csdn.net/chosen0ne/article/details/40782991)
- [vim 安装vim-go 打造GOLANG 专用IDE](http://studygolang.com/articles/2927)
- [为Vim配置Golang开发环境](http://ju.outofmemory.cn/entry/89207)

##### Atom
>- [打造atom成为golang开发神器](http://blog.csdn.net/sweetvvck/article/details/50333327)
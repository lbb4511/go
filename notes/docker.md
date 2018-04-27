## Docker常用命令

1. 查看docker信息（version、info）
    - 查看docker版本
    `$docker version`
    - 显示docker系统的信息
    `$docker info`
2. 对image的操作（search、pull、images、rmi、history）
    - 检索image
    `$docker search image_name`
    - 下载image
    `$docker pull image_name`
    - 列出镜像列表; -a, --all=false 显示所有镜像; --no-trunc=false 不要截断输出; -q, --quiet=false 仅显示数字ID
    `$docker images`
    - 删除一个或者多个镜像; -f, --force=false 强行; --no-prune=false 不要删除未标记的父级
    `$docker rmi image_name`
    - 显示一个镜像的历史; --no-trunc=false 不要截断输出; -q, --quiet=false 仅显示数字ID
    `$docker history image_name`
3. 启动容器（run）

    docker容器可以理解为在沙盒中运行的进程。这个沙盒包含了该进程运行所必须的资源，包括文件系统、系统类库、shell 环境等等。但这个沙盒默认是不会运行任何程序的。你需要在沙盒中运行一个进程来启动某一个容器。这个进程是该容器的唯一进程，所以当该进程结束的时候，容器也会完全的停止。

    - 在容器中运行"echo"命令，输出"hello word"
    `$docker run image_name echo "hello word"`
    - 交互式进入容器中
    `$docker run -i -t image_name /bin/bash`
    - 挂载本地目录,实现文件共享
    `$docker run -it -v _path:docker_path image_name /bin/bash`
    - 在容器中安装新的程序
    `$docker run image_name apt-get install -y app_name`

    *注意：  在执行apt-get 命令的时候，要带上-y参数。如果不指定-y参数的话，apt-get命令会进入交互模式，需要用户输入命令来进行确认，但在docker环境中是无法响应这种交互的。apt-get 命令执行完毕之后，容器就会停止，但对容器的改动不会丢失。*
4. 查看容器（ps）
    - 列出当前所有正在运行的container
    `$docker ps`
    - 列出所有的container$docker ps -a
    - 列出最近一次启动的container
    `$docker ps -l`
5. 保存对容器的修改（commit）

    当你对某一个容器做了修改之后（通过在容器中运行某一个命令），可以把对容器的修改保存下来，这样下次可以从保存后的最新状态运行该容器。

    - 保存对容器的修改; -a, --author="" 作者; -m, --message="" 提交消息
    `$docker commit ID new_image_name`

    *注意：  image相当于类，container相当于实例，不过可以动态给实例安装新软件，然后把这个container用commit命令固化成一个image。*
 6. 对容器的操作（rm、stop、start、kill、logs、diff、top、cp、restart、attach）
    - 删除所有容器
    ```
    $docker rm `docker ps -a -q`
    ```
    - 删除单个容器; -f, --force=false; -l, --link=false 删除指定的链接，而不是基础容器; -v, --volumes=false 删除与容器关联的卷
    `$docker rm Name/ID`
    - 停止、启动、杀死一个容器
    `$docker stop Name/ID`
    `$docker start Name/ID`
    `$docker kill Name/ID`
    - 从一个容器中取日志; -f, --follow=false 按照日志输出; -t, --timestamps=false 显示时间戳
    `$docker logs Name/ID`
    - 列出一个容器里面被改变的文件或者目录，list列表会显示出三种事件，A 增加的，D 删除的，C 被改变的
    `$docker diff Name/ID`
    - 显示一个运行的容器里面的进程信息
    `$docker top Name/ID`
    - 从容器里面拷贝文件/目录到本地一个路径
    `$docker cp Name:/container_path to_path`
    `$docker cp ID:/container_path to_path`
    - 重启一个正在运行的容器; -t, --time=10 尝试在停止容器之前停止的秒数, 默认=10
    `$docker restart Name/ID`
    - 附加到一个运行的容器上面; --no-stdin=false 不要附加标准输入(stdin); --sig-proxy=true 将所有接收到的信号代理给进程
    `$docker attach ID`

    *注意： attach命令允许你查看或者影响一个运行的容器。你可以在同一时间attach同一个容器。你也可以从一个容器中脱离出来，是从CTRL-C。*
7. 保存和加载镜像（save、load）

    当需要把一台机器上的镜像迁移到另一台机器的时候，需要保存镜像与加载镜像。

    - 保存镜像到一个tar包; -o, --output="" 写入文件
    `$docker save image_name -o file_path`
    - 加载一个tar包格式的镜像; -i, --input="" 从tar存档文件读取
    `$docker load -i file_path`
    - 机器a
    `$docker save image_name > /home/save.tar`
    - 使用scp将save.tar拷到机器b上，然后：
    `$docker load < /home/save.tar`
8. 登录registry server（login）
    - 登陆registry server; -e, --email="" 电子邮件; -p, --password="" 密码; -u, --username="" 用户名
    `$docker login`
 9. 发布image（push）
    - 发布docker镜像
    `$docker push new_image_name`
10. 根据Dockerfile 构建出一个容器
    - build 
    --no-cache=false 在构建映像时不要使用缓存
    -q, --quiet=false 抑制由容器生成的详细输出
    --rm=true 成功构建后删除中间容器
    -t, --tag="" 在成功的情况下应用于生成的图像的存储库名称（以及可选的标记）
    `$docker build -t image_name Dockerfile_path`

## Docker源码分析
- init函数的执行。在Golang中init函数的特性如下：
    - init函数用于程序执行前包的初始化工作，比如初始化变量等；
    - 每个包可以有多个init函数；
    - 包的每一个源文件也可以有多个init函数；
    - 同一个包内的init函数的执行顺序没有明确的定义；
    - 不同包的init函数按照包导入的依赖关系决定初始化的顺序；
    - init函数不能被调用，而是在main函数调用前自动被调用。
- os.Setenv("DEBUG", "1") 创建系统环境变量
- [Docker-github](https://github.com/docker/docker)
- Docker源码分析 孙宏亮 -infoQ
    - [Docker源码分析（一）：Docker架构](http://www.infoq.com/cn/articles/docker-source-code-analysis-part1)
    - [Docker源码分析（二）：Docker Client创建与命令执行](http://www.infoq.com/cn/articles/docker-source-code-analysis-part2)
    - [Docker源码分析（三）：Docker Daemon启动](http://www.infoq.com/cn/articles/docker-source-code-analysis-part3)
    - [Docker源码分析（四）：Docker Daemon之NewDaemon实现](http://www.infoq.com/cn/articles/docker-source-code-analysis-part4)
    - [Docker源码分析（五）：Docker Server的创建](http://www.infoq.com/cn/articles/docker-source-code-analysis-part5)
    - [Docker源码分析（六）：Docker Daemon网络](http://www.infoq.com/cn/articles/docker-source-code-analysis-part6)
    - [Docker源码分析（七）：Docker Container网络 （上）](http://www.infoq.com/cn/articles/docker-source-code-analysis-part7)
    - [Docker源码分析（八）：Docker Container网络 （下）](http://www.infoq.com/cn/articles/docker-source-code-analysis-part8)
    - [Docker源码分析（九）：Docker镜像](http://www.infoq.com/cn/articles/docker-source-code-analysis-part9)
    - [Docker源码分析（十）：Docker镜像下载](http://www.infoq.com/cn/articles/docker-source-code-analysis-part10)
    - [Docker源码分析（十一）：镜像存储](http://www.infoq.com/cn/articles/docker-source-code-analysis-part11)
    - [Docker三年回顾：梦想依在，人生正当年](http://www.infoq.com/cn/articles/docker-turns-3)
    - [超大规模容器调度系统的设计与实现](http://www.infoq.com/cn/presentations/how-to-design-an-ultra-large-scale-container-scheduling-system)

## 图书
- [Docker入门实战](https://yuedu.baidu.com/ebook/d817967416fc700abb68fca1?pn=1&rf=http%3A%2F%2Fyuedu.baidu.com%2Febook%2Fd817967416fc700abb68fca1%3Ffr%3Daladdin%26key%3Ddocker)
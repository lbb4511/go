# 项目列表

看过awesome-go项目，汇总了很多go开源项目。但是awesome-go收集了太全了，而且每个项目没有详细描述。因此我自己根据go语言中文社区提供的资料，还有互联网企业架构设计中的常见组件分类，共精心挑选了100多个开源项目（项目不限于在github开源的项目），分成以下十几个大类。这个项目可以理解为互联网IT人打造的中文版awesome-go。这个项目初衷是帮助到那些想学习和借鉴优秀golang开源项目，和在互联网架构设计时期望快速寻找合适轮子的人。

## 监控系统

- [OpenFalcon](https://github.com/open-falcon/of-release)是一款小米开源的监控系统。功能：数据采集免配置：agent自发现、支持Plugin、主动推送模式;容量水平扩展：生产环境每秒50万次数据收集、告警、存储、绘图，可持续水平扩展。告警策略自发现：Web界面、支持策略模板、模板继承和覆盖、多种告警方式、支持回调动作。告警设置人性化：支持最大告警次数、告警级别设置、告警恢复通知、告警暂停、不同时段不同阈值、支持维护周期，支持告警合并。历史数据高效查询：秒级返回上百个指标一年的历史数据。Dashboard人性化：多维度的数据展示，用户自定义Dashboard等功能。架构设计高可用：整个系统无核心单点，易运维，易部署。
- [banshee](https://github.com/eleme/banshee)周期性指标的监控系统.
- [Kapacitor](https://github.com/influxdata/kapacitor)是一个开源框架，用来处理、监控和警告时间序列数据。
- [Pome](https://github.com/rach/pome)是PostgresMetrics的意思。Pome是一个PostgreSQL的指标仪表器，用来跟踪你的数据库的健康状况。
- [pingd](https://github.com/pinggg/pingd)是世界上最简单的监控服务，使用golang编写。软件支持IPv6，但是服务器不支持.pingd允许同时ping上千个IPs，在此期间还可以管理监控的主机。用户提供主机名或者IP，还有用户邮箱地址，就可以使用3个生成URLs来开启，停止或者删除你的追踪。每当你的服务器停机或者后台在线都会发送通知，还包含控制URLs。
- [actiontechzabbixmysqlmonitor](https://github.com/actiontech/actiontech_zabbix_mysql_monitor)是perconamonitoringpluginszabbix 的Go语言版本，是由爱可生公司开源的MySQL监控插件和模板，整合上百个性能监控指标，支持LowLevelDiscovery自动发现多实例环境，支持performance_schema
- [rtop](https://github.com/rapidloop/rtop)top是一个简单的无代理的远程服务器监控工具，基于SSH连接进行工作。无需在被监控的服务器上安装任何软件。rtop直接通过SSH连接到待监控服务器，然后执行命令来收集监控数据。rtop每几秒钟就自动更新监控数据，类似其他*top命令
- [Prometheus](https://github.com/prometheus/prometheus)Prometheus是一个开源的服务监控系统和时间序列数据库，提供监控数据存储，展示，告警等功能
- [bosun](https://github.com/bosun-monitor/bosun)专业的跨平台开源系统监控项目，go语言编写，灵活的模板和表达式配合上各种collector可以监控任何应用或系统级的运行数据，比zabbix更轻量级、更易入手和更适合定制。
- [urlooker](https://github.com/urlooker)监控web服务可用性及访问质量，采用go语言编写，易于安装和二次开发.支持一下特性：返回状态码检测;页面响应时间检测;页面关键词匹配检测;带cookie访问;agent多机房部署，指定机房访问;检测结果支持向open-falcon推送;支持短信和邮件告警
- [satellite](https://github.com/gravitational/satellite)用于监测kubernetes健康状态的一个工具／库。其特点是：轻量级定期测试，高可用性和弹性网络分区，无单点故障，以时间序列的格式存储监控数据。
- [checkup](https://github.com/sourcegraph/checkup)一个分布式的无锁的站点健康状态检查工具。支持检查http，tcp，dns等的状态并可将结果保存在s3。自带了一个美观的界面。
- [zabbixctl](https://github.com/kovetskiy/zabbixctl)Zabbixctl是采用Zabbix服务API的命令行工具，它提供了有效的方式去查询和处理trigger状态、主机最新数据和用户组。
- [cloudinsight-agent](https://github.com/cloudinsight/cloudinsight-agent/)提供可视化监控的saas平台cloudinsight开源的一个监控客户端。Cloudinsight探针可以收集它所在操作系统的各种指标，然后发送到Cloudinsight后端服务
- [owl](https://github.com/TalkingData/owl)是TalkingData公司推出的一款开源分布式监控系统,演示环境`http://54.223.127.87/`，登录账号密码`demo/demo`
- [SmartPing](https://github.com/gy-games/smartping)SmartPing为一个各机器(点)间间互PING检测工具，支持互PING，单向PING，绘制拓扑及报警功能。系统设计为无中心化原则，所有的数据均存储自身点中，默认数据循环保留1个月时间，由自身点的数据绘制出PING包的状态，由各其他点的数据绘制进PING包的状态，并API接口获取其他点数据绘制整体PING拓扑图，拓扑图中存在报警功能。

## 容器技术

- [Pouch](https://github.com/alibaba/pouch)是Alibaba公司开源的容器引擎技术，其主要功能包括基本的容器管理能力，安全稳定的强容器隔离能力，以及对应用无侵入性的富容器技术。
- [SwarmKit](https://github.com/docker/swarmkit)是Docker公司开源的Docker集群管理和容器编排工具，其主要功能包括节点发现、基于raft算法的一致性和任务调度等。
- [DaoliNet](https://github.com/daolinet/daolinet)是一个软件定义网络(SDN)系统，其设计目的是为Docker容器提供动态、高效的链接。在Docker容器中，微服务工作负载具有轻量且短暂的性质，DaoliNet恰好适用于这种性质。
- [Harbor](https://github.com/vmware/harbor)容器应用的开发和运行离不开可靠的镜像管理。从安全和效率等方面考虑，部署在私有环境内的Registry是非常必要的。ProjectHarbor是由VMware公司中国团队为企业用户设计的Registryserver开源项目，包括了权限管理(RBAC)、LDAP、审计、管理界面、自我注册、HA等企业必需的功能，同时针对中国用户的特点，设计镜像复制和中文支持等功能
- [REX-Ray](https://github.com/emccode/rexray)是一个EMC{code}团队领导的开源项目，为Docker、Mesos及其他容器运行环境提供持续的存储访问。其设计旨在囊括通用存储、虚拟化和云平台，提供高级的存储功能。
- [Clair](https://github.com/coreos/clair)是一个容器漏洞分析服务。它提供一个能威胁容器漏洞的列表，并且在有新的容器漏洞发布出来后会发送通知给用户。
- [Weave](https://github.com/zettio/weave)创建一个虚拟网络并连接到部署在多个主机上的Docker容器。
- [Rocket](https://github.com/coreos/rkt)（也叫rkt）是CoreOS推出的一款容器引擎，和Docker类似，帮助开发者打包应用和依赖包到可移植容器中，简化搭环境等部署工作。Rocket和Docker不同的地方在于，Rocket没有Docker那些为企业用户提供的“友好功能”，比如云服务加速工具、集群系统等。反过来说，Rocket想做的，是一个更纯粹的业界标准。
- [libnetwork](https://github.com/docker/libnetwork)提供一个原生Go实现的容器连接，是容器的网络。libnetwork的目标是定义一个健壮的容器网络模型（ContainerNetworkModel），提供一个一致的编程接口和应用程序的网络抽象。
- [Wormhole](https://github.com/vishvananda/wormhole)是一个能识别命名空间的由Socket激活的隧道代理。可以让你安全的连接在不同物理机器上的Docker容器。可以用来完成一些有趣的功能，例如连接运行在容器本机的服务或者在连接后创建按需的服务。
- [Shipyard](https://github.com/shipyard/shipyard)是一个基于Web的Docker管理工具，支持多host，可以把多个Dockerhost上的containers统一管理；可以查看images，甚至buildimages；并提供RESTfulAPI等等。Shipyard要管理和控制Dockerhost的话需要先修改Dockerhost上的默认配置使其支持远程管理。
- [Docker](https://github.com/docker/docker)是一个开源的应用容器引擎，让开发者可以打包他们的应用以及依赖包到一个可移植的容器中，然后发布到任何流行的Linux机器上，也可以实现虚拟化。容器是完全使用沙箱机制，相互之间不会有任何接口（类似iPhone的app）。几乎没有性能开销,可以很容易地在机器和数据中心中运行。最重要的是,他们不依赖于任何语言、框架或包装系统。
- [scope](https://github.com/weaveworks/scope)一个docker&kubernetes的管理，监控可视化工具，可以看到容器间的拓扑关系和tcp通信
- [habitus](https://github.com/cloud66/habitus)一个快速实现dockerbuild流程的工具，支持复杂的dockerbuild流程，实现多个dockerfile的build流程，典型应用如将需要静态编译的程序，如go，java这类程序在一个dockerbuild编译好之后，得到的二进制包用到后续的build流程
- [sextant](https://github.com/k8sp/sextant)提供了可以通过PXE全自动化安装初始化一个CoreOS+kubernetes集群。
- [KubeVirt](http://github.com/kubevirt/kubevirt)是一个围绕Kubernetes构建的虚拟机管理架构。使用KubeVirt可声明：创建预定义的VM，在Kubernetes集群上调度VM（启动VM，停止VM，删除VM）

## PaaS工具

- [Kel](https://github.com/kelproject)是一个开源的基于Kubernetes构建的PaaS系统，采用Python和Go语言开发。Kel可简化管理Web应用发布和托管整个软件生命周期。Kel帮助开发和运维人员轻松管理他们的应用架构，通过一组工具和组件让K8S使用非常简单。
- [CloudFoundry-Mesos](https://github.com/mesos/cloudfoundry-mesos)框架由华为与Mesosphere的工程师合作完成，能够为应用提供安全可靠的、可伸缩、可扩展的云端运行环境，并且应用能够享用CloudFoundry生态圈内各类丰富的服务资源。企业能够通过CloudFoundry开发云应用，并通过CloudFoundry-Mesos将应用部署到DCOS上，使应用能够与DCOS上安装的其他服务及应用框架共享资源，实现资源利用率最大化，能够大幅降低企业数据中心运营成本。DCOS能够运行在虚拟和物理环境上，能够支持Linux（以及很快支持Windows），并可适用于私有云、公有云及混合云环境。
- [Flynn](https://github.com/github/flynn)是一个开源的PaaS系统，由Docker开发。采用Go语言编写。支持数据库包括Postgres、Redis和MongoDB.Flynn使用完全组件化模块化的设计，任何一个组件和模块都可以独立的进行替换。
- [DINP](https://git.oschina.net/cnperl/dinp-server)是又一个基于Docker开发的PaaS平台。
- [Kubernetes](https://github.com/kubernetes/kubernetes)是来自Google云平台的开源容器集群管理系统。基于Docker构建一个容器的调度服务。该系统可以自动在一个容器集群中选择一个工作容器供使用。其核心概念是ContainerPod。
- [Tsuru](https://github.com/tsuru/tsuru)在Tsuru的PaaS服务下，你可以选择自己的编程语言，选择使用SQL或者NoSQL数据库，memcache、redis、等等许多服务，甚至与你可以使用Git版本控制工具来上传你应用。
- [atlantis](https://github.com/ooyala/atlantis)是一款基于Docker，使用Go编写，为HTTP应用准备的开源PaaS。Atlantis可以在路由请求中轻松的构建和部署应用到容器。Atlantis在Ooyala的新应用中得到了很广泛的应用。
- [lain](https://github.com/laincloud/lain)是一个基于docker的PaaS系统。其面向技术栈多样寻求高效运维方案的高速发展中的组织，devops人力缺乏的startup，个人开发者。统一高效的开发工作流，降低应用运维复杂度；在IaaS/私有IDC裸机的基础上直接提供应用开发，集成，部署，运维的一揽子解决方案。
- [OpenDCP](https://github.com/weibocom/opendcp)是一个基于Docker的云资源管理与调度平台，集镜像仓库、多云支持、服务编排、服务发现等功能与一身，支持服务池的扩缩容，其技术体系源于微博用于支持节假日及热点峰值流量的弹性调度DCP系统。OpenDCP允许利用公有云服务器搭建起适应互联网应用的IT基础设施，并且将运维的工作量降到最低。
- [Swan](http://github.com/Dataman-Cloud/swan)是一个基于mesos的新HTTPAPI，使用golang开发。你可以使用swan在mesos集群上部署应用程序，并管理应用程序的整个生命周期，还可以使用新版本进行滚动更新，扩展应用程序，并且可以在应用程序或服务不可用时对应用程序和自动故障转移进行运行状况检查。
- [Deis](https://github.com/deis/deis)是一个基于Docker和CoreOS的开源PaaS平台，旨在让部属和管理服务器上的应用变得轻松容易。它可以运行在AWS、GCE以及Openstack平台下。详情，可了解《Deisv1.0正式发布！》。

## 大数据&ML

- [MLF](https://github.com/huichen/mlf)弥勒佛项目是一个大数据机器学习框架。具有为处理大数据优化，可随业务增长scaleup，模型的训练和使用都可以作为library或者service整合到在生产系统中，具有丰富的模型，高度可扩展，高度可读性，适合初学者进行大数据模型的学习等特点
- [Glow](https://github.com/chrislusf/glow)是使用Go编写的易用分布式计算系统，是HadoopMapReduce，Spark，Flint，Samza等等的替代品。Glow的目标是提供一个库，可以在并行线程或者分布式集群机器中进行更简单计算。
- [goml](https://github.com/cdipaolo/goml)机器学习的库,包含了许多工具，能让你以在线方式学习其频道的数据内容。
- [Golearn](https://github.com/sjwhitworth/golearn)实现了熟悉的Scikit-learn适应/预测界面，可实现快速预估测试和交换。GoLearn是一个成熟的项目，它提供了交叉验证和训练/测试等辅助功能。
- [Gorgonia](https://github.com/chewxy/gorgonia)这个机器学习资料库完全是用Go语言编写而成，据其开发者“chewxy”称能“提供动态建立神经网络及相关算法必需条件。”
- [poseidon](https://github.com/Qihoo360/poseidon)系统是一个日志搜索平台，可以在数百万亿条、数百PB大小的日志数据中快速分析和检索特定字符串。该系统可以应用于任何结构化或非结构化海量(从万亿到千万亿规模)数据的查询检索需求)。

## 微服务

- [micro](https://github.com/micro/micro)是一个专注于简化分布式系统开发的微服务生态系统。可插拔的插件化设计，提供强大的可插拔的架构来保证基础组件可以被灵活替换。
- [kite](https://github.com/koding/kite)一个基于go语言的微服务框架,Kite是Koding公司内部的一个框架,该框架提供服务发现，多种认证功能，服务端通过RPC进行通信，同时还提供了websocket的js库，方便浏览器于服务器间进行通信。
- [goa](https://github.com/goadesign/goa)是一款用Go用于构建微服务的框架，采用独特的设计优先的方法。
- [Go-kit](https://github.com/go-kit/kit)是一个Go语言的分布式开发包，用于开发微服务。
- [Hprose](https://github.com/andot/hprose)是高性能远程对象服务引擎（HighPerformanceRemoteObjectServiceEngine）的缩写——微服务首选引擎。它是一个先进的轻量级的跨语言跨平台面向对象的高性能远程动态通讯中间件。它不仅简单易用，而且功能强大。你只需要稍许的时间去学习，就能用它轻松构建跨语言跨平台的分布式应用系统了。
- [Gizmo](https://github.com/NYTimes/gizmo)纽约时报开源的go微服务工具.提供如下特性:标准化配置和日志;可配置策略的状态监测端点;用于管理pprof端点和日志级别的配置;结构化日志，提供基本请求信息;端点的有用度量;优雅的停止服务;定义期待和词汇的基本接口
- [hystrix-go](https://github.com/afex/hystrix-go)用来隔离远程系统调用，第三方库调用，服务调用，提供熔断机制，避免雪崩效应的库，Hystrix的go版本。注Hystrixs是Netflix开源的一个java库
- [gateway](https://github.com/fagongzi/gateway)是一个使用go实现的基于HTTP的API网关。**特性**：API聚合;流控;熔断;负载均衡;健康检查;监控;消息路由;后端管理WebUI.**能做什么**：规划更友好的URL给调用者。聚合多个API的结果返回给API调用者，利于移动端，后端可以实现原子接口。保护后端API服务不会被突发异常流量压垮。提供熔断机制，使得后端APIServer具备自我恢复能力。借助消息路由能力，实现灰度发布，AB测试。
- [fabio](https://github.com/eBay/fabio)是ebay团队用golang开发的一个快速、简单零配置能够让consul部署的应用快速支持http(s)的负载均衡路由器。这里有一篇中文文章<http://dockone.io/article/1567>介绍了如何用fabio＋consul实现服务发现，负载均衡，并阐述了原理，最后还有demo程序
- [appdash](https://github.com/sourcegraph/appdash)go版本的分布式应用跟踪系统，基于googledapper的原理构建
- [Jaeger](https://github.com/uber/jaeger)是Uber的分布式跟踪系统，基于googledapper的原理构建，以Cassandra作为存储层

## CI/CD

- [Cyclone](https://github.com/caicloud/cyclone)Cyclone是一个打造容器工作流的云原生持续集成持续发布平台，简单易用，使用Go语言开发，有详尽的中文文档
- [Drone](https://github.com/drone/drone)Drone是一个基于Docker的持续发布平台，使用Go语言开发

## 数据库技术

- [LedisDB](https://github.com/siddontang/ledisdb)ledisdb是一个参考ssdb，采用go实现，底层基于leveldb，类似redis的高性能nosql数据库，提供了kv，list，hash以及zset数据结构的支持。
- [BuntDB](github.com/tidwall/buntdb)是纯Go开发的、低层级的（low-level）的、可嵌入的key/value内存数据库（IMDB），数据持久化存储，遵从ACID，支持自定义索引和geospatial数据。
- [Cockroach](https://github.com/cockroachdb/cockroach)CockroachDB(蟑螂数据库）是一个可伸缩的、支持地理位置处理、支持事务处理的数据存储系统。CockroachDB提供两种不同的的事务特性，包括快照隔离（snapshotisolation，简称SI）和顺序的快照隔离（SSI）语义，后者是默认的隔离级别。
- [qb-go](https://github.com/aacanakin)qb是用来让使更容易使用数据库的go语言的数据库工具包。它受Python最喜欢的ORMSQLAlchemy的启发，既是一个ORM，也是一个查询生成器。它在表达api和查询构建东西的情形下是相当模块化的。
- [GoshawkDB](https://github.com/goshawkdb)GoshawkDB是一个采用Go语言开发支持多平台的分布式的对象存储服务，支持事务以及容错。GoshawkDB的事务控制是在客户端完成的。GoshawkDB服务器端使用AGPL许可，而Go语言客户端使用Apache许可证
- [Codis](https://github.com/wandoulabs/codis)Codis是一个分布式Redis解决方案,对于上层的应用来说,连接到CodisProxy和连接原生的RedisServer没有明显的区别(不支持的命令列表),上层应用可以像使用单机的Redis一样使用,Codis底层会处理请求的转发,不停机的数据迁移等工作,所有后边的一切事情,对于前面的客户端来说是透明的,可以简单的认为后边连接的是一个内存无限大的Redis服务.
- [Cayley](https://github.com/google/cayley)Cayley是Google的一个开源图(Graph)数据库，其灵感来自于Freebase和Google的KnowledgeGraph背后的图数据库。
- [Redigo](https://github.com/garyburd/redigo)Redigo是Redis数据库的Go客户端。
- [radix.v2](https://github.com/mediocregopher/radix.v2)radix.v2是redis官方推荐的客户端之一，相比于redigo,radix.v2特点是轻量、接口实现优雅、API友好
- [redis-go-cluster](https://github.com/chasex/redis-go-cluster)redis-go-cluster是基于Redigo实现的GolangRedis客户端。redis-go-cluster可以在本地缓存slot信息，并且当集群修改的时候会自动更新。此客户端管理每个节点连接池，使用goroutine来尽可能的并发执行，达到了高效，低延迟。
- [elastic](https://github.com/olivere/elastic)elastic是开源搜索引擎elasticsearch的golang客户端，API友好，支持绝大部分es的接口,支持的es版本全面，从1.x到最新的6.x全覆盖
- [Dgraph](https://github.com/dgraph-io/dgraph)dgraph是可扩展的，分布式的，低延迟图形数据库。DGraph的目标是提供Google生产水平的规模和吞吐量，在超过TB的结构数据里，未用户提供足够低延迟的实时查询。DGraph支持GraphQL作为查询语言，响应JSON。
- [DegDB](https://github.com/degdb/degdb)DegDB是分布式的经济图数据库。
- [Vitess](https://github.com/youtube/vitess)outube出品的开源分布式MySQL工具集Vitess，自动分片存储MySQL数据表，将单个SQL查询改写为分布式发送到多个MySQLServer上，支持行缓存（比MySQL本身缓存效率高），支持复制容错，已用于Youtube生产环境
- [xuncache](https://github.com/sun8911879/xuncache)xuncache是免费开源的NOSQL(内存数据库)采用golang开发,简单易用而且功能强大(就算新手也完全胜任)、性能卓越能轻松处理海量数据,可用于缓存系统.
- [pgweb](https://github.com/sosedoff/pgweb)gweb是一个采用Go语言开发的基于Web的PostgreSQL管理系统。
- [Orchestrator](https://github.com/outbrain/orchestrator)MySQL复制拓扑可视化工具
- [mysql-schema-sync](https://github.com/hidu/mysql-schema-sync)mysql-schema-sync是一款使用go开发的、跨平台的、绿色无依赖的MySQL表结构自动同步工具。用于将线上(其他环境)数据库结构变化同步到测试（本地）环境!
- [TiDB](https://github.com/pingcap/tidb)TiDB是国内PingCAP团队开发的一个分布式SQL数据库。其灵感来自于Google的F1,TiDB支持包括传统RDBMS和NoSQL的特性。
- [kingshard](https://github.com/flike/kingshard)一个高性能的mysql中间件，支持读写分离，数据分片，安全审计等功能
- [influxdb](https://github.com/influxdata/influxdb)一个可以水平扩展的时间序列数据库，内建httpapi，支持对数据打tag，灵活的查询策略和数据的实时查询，支持类sql语句进行查询
- [DBShield](http://github.com/nim4/DBShield)DBShield是用Go语言开发的数据库防火墙，用来保护你的数据免受SQL注入的侵扰。支持的数据库包括：DB2、MySQL、MariaDB、Oracle、PostgreSQL。

## 存储技术

- [Torus](https://github.com/coreos/torus)Torus是一种针对容器集群量身打造的存储系统，可以为通过Kubernetes编排和管理的容器集群提供可靠可扩展的存储。这是继etcd、rkt、flannel，以及CoreOSLinux之后CoreOS发布的另一个开源产品。
- [Afero](https://github.com/spf13/afero)Afero是一个文件系统框架，提供一个简单、统一和通用的API和任何文件系统进行交互，作为抽象层还提供了界面、类型和方法。Afero的界面十分简洁，设计简单，舍弃了不必要的构造函数和初始化方法。Afero作为一个库还提供了一组可交互操作的后台文件系统，这样在与Afero协作时，还可以保留os和ioutil软件包的功能和好处。
- [REX-Ray](https://github.com/emccode/rexray)REX-Ray是一个EMC{code}团队领导的开源项目，为Docker、Mesos及其他容器运行环境提供持续的存储访问。其设计旨在囊括通用存储、虚拟化和云平台，提供高级的存储功能。
- [SeaweedFS](https://github.com/chrislusf/seaweedfs)SeaweedFS是简单，高伸缩性的分布式文件系统，包含两部分：存储数十亿的文件；快速为文件服务。SeaweedFS作为支持全POSIX文件系统语义替代，Seaweed-FS选择仅实现key-file的映射，类似"NoSQL"，也可以说是"NoFS"。
- [bfs](https://github.com/Terry-Mao/bfs)bfs是使用Go编写的分布式文件系统（小文件存储）。
- [IPFS](https://github.com/ipfs/go-ipfs)IPFS是分布式文件系统，寻求连接所有计算机设备的相同文件系统。在某些方面，这很类似于原始的Web目标，但是IPFS最终会更像单个比特流群交换的git对象。IPFS＝InterPlanetaryFileSystem
- [Gotgt](https://github.com/gostor/gotgt)Gotgt是使用Go编写的高性能、可扩展的iSCSItarget服务。
- [Etcd](https://github.com/coreos/etcd)&Fleetetcd是由CoreOS开发并维护键值存储系统，它使用Go语言编写，并通过Raft一致性算法处理日志复制以保证强一致性。目前，Google的容器集群管理系统Kubernetes、开源PaaS平台CloudFoundry和CoreOS的Fleet都广泛使用了etcd。详情，可了解《Etcd：用于服务发现的键值存储系统》。Fleet则是一个分布式的初始化系统。它们之所以选择使用Go语言，则是因为Go语言对跨平台的良好支持，以及其背后的强大社区。
- [Syncthing](https://github.com/syncthing/syncthing)一款用Go语言编写的开源云存储和同步服务工具，用户的数据将由自己完全控制，所有的通信全都加密，每个访问节点都用加密证书验证。该项目被认为是Dropbox和BitTorrentSync的开源替代，详情。Syncthing之所以选择Go语言，也是出于跨平台考虑。
- [cache2go](https://github.com/muesli/cache2go)比较简单的一个缓存库，代码量很少，适合新手学习，可以学习到锁、goroutines等。
- [groupcache](https://github.com/golang/groupcache)与memcached同一作者，相当于是memcached的go语言实现。

## 分布式系统

- [Confd](https://github.com/kelseyhightower)Confd是一个轻量级的配置管理工具。通过查询Etcd，结合配置模板引擎，保持本地配置最新，同时具备定期探测机制，配置变更自动reload。
- [zerg](https://github.com/huichen/zerg)基于docker的分布式爬虫服务
- [Doorman](https://github.com/youtube/doorman)Doorman是一个客户端速率限制的解决方案，客户端与共享资源进行通讯，包括数据库、gRPC服务、RESTfulAPI等等可使用Doorman来限制对资源的调用。Doorman使用Go语言开发，使用gRPC的通讯协议。其高可用特性需要一个分布式的锁管理器，当前支持etcd，也可使用Zookeeper替代。
- [mgmt](https://github.com/purpleidea/mgmt)mgmt是一个分布式的，事件驱动的配置管理工具。该工具支持并行执行，其librarification作为新的及已存在的软件的基础管理工具。
- [Yoke](https://github.com/nanopack/yoke)Yoke是Postgres的高可用集群，具有自动切换和自动集群恢复。Postgres冗余/自动故障转移解决方案，提供一个高可用PostgreSQL集群的简单管理。
- [SeaweedFS](https://github.com/chrislusf/seaweedfs)SeaweedFS是简单，高伸缩性的分布式文件系统，包含两部分：存储数十亿的文件；快速为文件服务。SeaweedFS作为支持全POSIX文件系统语义替代，Seaweed-FS选择仅实现key-file的映射，类似"NoSQL"，也可以说是"NoFS"。
- [Glow](https://github.com/chrislusf/glow)Glow是使用Go编写的易用分布式计算系统，是HadoopMapReduce，Spark，Flint，Samza等等的替代品。Glow的目标是提供一个库，可以在并行线程或者分布式集群机器中进行更简单计算。
- [Nomad](https://github.com/hashicorp/nomad)Nomad是一个集群管理器和调度器，专为微服务和批量处理工作流设计。Nomad是分布式，高可用，可扩展到跨数据中心和区域的数千个节点。
- [dcmp](https://github.com/silenceper/dcmp)DCMP是分布式配置管理平台。提供了一个etcd的管理界面，可通过界面修改配置信息，借助confd可实现配置文件的同步。
- [gleam](https://github.com/chrislusf/gleam)此处是一个通过Go和LuaJIT编写的快速和可扩展的分布式map/reduce系统，很好的将Go的高并发性与Luajit高性能相结合，可独立运行或用于分布式计算。

## 消息系统

- [KiteQ](https://github.com/blackbeans/kiteq)KiteQ是一个基于go+protobuff实现的多种持久化方案的mq框架（消息队列）。
- [NSQ](https://github.com/bitly/nsq)NSQ是无中心设计、节点自动注册和发现的开源消息系统。可作为内部通讯框架的基础，易于配置和发布。
- [kingtask](https://github.com/kingsoft-wps/kingtask)kingtask是一个由Go开发的轻量级的异步定时任务系统。支持定时的异步任务。支持失败重试机制，重试时刻和次数可自定义。任务执行结果可查询。
- [GoMachinery](https://github.com/RichardKnop/machinery)Machinery是一个Go语言的异步任务队列和作业队列，基于分布式消息传递。类似Python的Celery框架。
- [kaca](https://github.com/scottkiss/kaca)kaca是用golang语言开发的基于websocket协议的消息发布/订阅系统。

## 服务器管理

- [Sharkey](https://github.com/square/sharkey)Sharkey是OpenSSH管理证书使用的服务。Sharkey分为客户端组件和服务端组件，服务端负责发布已签署的主机证书，客户端负责在机器上安装主机证书。
- [OSinstall](https://github.com/idcos/osinstall)CloudBoot”(OSinstall)云装机平台，是金融云初创公司杭州云霁科技推出的一款X86服务器全自动装机工具，遵循Apache协议，完全开源免费。全自动构建物理机资源池，像创建虚拟机一样方便的安装物理机。
- [ssh2go](https://github.com/karfield/ssh2go)ssh2go是对libssh的golang封装。libssh是SSH的代码库，同时支持服务端和客户端，日常所见的ssh,sshd,scp,sftp均基于libssh。ssh2go是对libssh的Go语言绑定，100%的libssh接口都可用，同时集成示例，方便参考。
- [Gooverssh](https://github.com/scottkiss/gooverssh)gooverssh是基于gosshtool的一个ssh开发包开发的一个基于ssh本地端口转发服务小应用，可以方便突破一些网络限制，如通过ssh代理访问内网数据库服
- [gosshtool](https://github.com/scottkiss/gosshtool)gosshtoolprovidesomeusefulfunctionsforsshclientingolang.implementedusinggolang.org/x/crypto/ssh.go语言中提供ssh相关操作，支持ssh本地端口转发服务
- [teleport](https://github.com/gravitational/teleport)teleport是初创公司Gravitational的一款基于ssh和https的远程管理linux集群服务器的工具，其特点是支持双重校验登陆；操作记录回放；session共享，便于协作排障；自动发现集群的服务器和容器

## 安全工具

- [gomitmproxy](https://github.com/sheepbao/gomitmproxy)GomitmProxy是想用golang语言实现的mitmproxy，主要实现http代理，目前实现了http代理和https抓包功能。
- [Hyperfox](https://github.com/xiam/hyperfox)Hyperfox是一个安全的工具用来代理和记录局域网中的HTTP和HTTPS通讯。
- [Gryffin](https://github.com/yahoo/gryffin)Gryffin是雅虎开发的一个大规模Web安全扫描平台。它不是另外一个扫描器，其主要目的是为了解决两个特定的问题——覆盖率和伸缩性。
- [ngrok](https://github.com/inconshreveable/ngrok)ngrok是一个反向代理，通过在公共的端点和本地运行的Web服务器之间建立一个安全的通道。ngrok可捕获和分析所有通道上的流量，便于后期分析和重放。

## 网络工具

- [DaoliNet](https://github.com/daolinet/daolinet)DaoliNet是一个软件定义网络(SDN)系统，其设计目的是为Docker容器提供动态、高效的链接。在Docker容器中，微服务工作负载具有轻量且短暂的性质，DaoliNet恰好适用于这种性质。
- [Seesaw](https://github.com/google/seesaw)Seesaw是Google开源的一个基于Linux的负载均衡系统。Seesaw包含基本的负载均衡特性，同时支持一些高级的功能，诸如：anycast,DirectServerReturn(DSR),支持多个VLANs和集中式配置。同时其设计的宗旨是易于维护。需要注意的是，尽管该项目挂靠在Google名下，但并非Google官方产品。
- [TcpRoute2](https://github.com/GameXG/TcpRoute2)TcpRoute,TCP层的路由器。对于TCP连接自动从多个线路(允许任意嵌套)、多个域名解析结果中选择最优线路。TcpRoute2是golang重写的版本。通过socks5代理协议对外提供服务。代理功能拆分成了独立的库，详细代理url格式级选项请参见ProxyClient，目前支持直连、socks4、socks4a、socks5、http、https、ss代理线路
- [Gor](https://github.com/buger/gor)Gor是用Go编写的简单HTTP流量复制工具，主要是为了从生产服务器返回流量到开发环境。使用Gor可以在实际的用户会话中测试代码。
- [Traefik](https://github.com/containous/traefik)Træfɪk是一个新型的http反向代理、负载均衡软件，能轻易的部署微服务.它支持多种后端(Docker,Swarm,Mesos/Marathon,Consul,Etcd,Zookeeper,BoltDB,RestAPI,file...),可以对配置进行自动化、动态的管理.
- [TChannel](https://github.com/uber/tchannel)TChannel是用于RPC的网络复用和成帧协议。
- [go-tcp-proxy](https://github.com/jpillora/go-tcp-proxy)go-tcp-proxy是一个简单的tcp代理，可以用于tcp端口转发，也可以用做http代理使用
- [myLG](https://github.com/mehrdadrad/mylg)myLG是一个开源的网络工具集，它包含了很多不同类型的网络诊断工具,功能包括ping，trace，bgp，dnslookup，端口扫描，局域网网络发现，提供web界面,tcpdump等
- [cow](https://github.com/cyfdecyf/cow)COW是一个简化穿墙的HTTP代理服务器。它能自动检测被墙网站，仅对这些网站使用二级代理；支持多种协议：sock5、http、shadow、cow

## Web工具

- [Tyk](https://github.com/lonelycode/tyk)Tyk是一个开源的、轻量级的、快速可伸缩的API网关，支持配额和速度限制，支持认证和数据分析，支持多用户多组织，提供全RESTfulAPI。
- [Shortme](https://github.com/andyxning/shortme)用Golang编写的URL短链接服务。
- [WuKongSearch](https://github.com/huichen/wukong)WuKong是一个全文搜索引擎。功能特性有：高效索引和搜索（1M条微博500M数据28秒索引完，1.65毫秒搜索响应时间，19K搜索QPS）；支持中文分词（使用sego分词包并发分词，速度27MB/秒）；支持计算关键词在文本中的紧邻距离（tokenproximity）；支持计算BM25相关度；支持自定义评分字段和评分规则；支持在线添加、删除索引；支持持久存储；可实现分布式索引和搜索等
- [Pholcus](https://github.com/henrylee2cn/pholcus)Pholcus（幽灵蛛）是一款纯Go语言编写的高并发、分布式、重量级爬虫软件，支持单机、服务端、客户端三种运行模式，拥有Web、GUI、命令行三种操作界面；规则简单灵活、批量任务并发、输出方式丰富（mysql/mongodb/csv/excel等）、有大量Demo共享；同时她还支持横纵向两种抓取模式，支持模拟登录和任务暂停、取消等一系列高级功能。
- [Codetainer](https://github.com/codetainerapp/codetainer)Codetainer可以让你创建基于浏览器上的代码运行沙箱，可方便的嵌入到你的Web应用中。你可以把它当成是codepicnic.com的开源克隆).
- [GoTTY](https://github.com/yudai/gotty)GoTTY是个简单的命令行工具，可以把CLI工具共享成Web应用。GoTTY可以把终端作为Web应用共享。
- [TermUI](https://github.com/gizak/termui)Go语言编写的终端仪表盘
- [Hound](https://github.com/etsy/hound)快如闪电的代码搜索开源工具
- [goim](https://github.com/Terry-Mao/goim)goim是一个支持集群的im及实时推送服务（支持websocket，http和tcp协议）
- [fasthttp](https://github.com/valyala/fasthttp)asthttp是Go的快速HTTP实现，当前在1M并发的生产环境使用非常成功，可以从单个服务器进行100Kqps的持续连接。HTTP服务器性能与net/http比较，fasthttp比net/http快10倍
- [netgraph](https://github.com/ga0/netgraph)netgraph是一个Go语言编写的跨平台的B/S架构的HTTP抓包工具，方便在Linux服务器上直接查看HTTP包。
- [gohttp](https://github.com/codeskyblue/gohttp)gohttp是一个http的文件服务器，功能有：各种文件的预览功能，实时的目录zip打包下载，二维码扫描下载的支持，苹果应用的在线安装，文件上传等
- [API-front](https://github.com/hidu/api-front)APIfront是HTTPAPI前端，可进行请求代理转发、协议抓包分析、流量复制。主要是用于开发测试环境,用来解决开发测试环境多变等问题
- [esumablefileuploads](https://github.com/tus)实现文件上传的断点续传功能，整套功能包含了协议实现，client，server。client及server有多种语言的实现包括go，python，node等｜
- [pproxy](https://github.com/hidu/pproxy)http抓包代理程序,http协议调试工具
- [hystrix-go](https://github.com/afex/hystrix-go)用来隔离远程系统调用，第三方库调用，服务调用，提供熔断机制，避免雪崩效应的库，Hystrix的go版本。注Hystrixs是Netflix开源的一个java库

## Web框架

- [Iris-Go](https://github.com/kataras/iris)通过Iris-Go，可以方便的帮助你来开发基于web的应用。简单来说：Iris-Go与国内大牛的BeeGo类似，但从其官方介绍的资料来看，Iris-Go的性能更优！
- [Baa](https://github.com/go-baa/baa)Baa一个简单高效的Goweb开发框架。主要有路由、中间件，依赖注入和HTTP上下文构成。
- [Orivil](https://github.com/orivil/orivil)Orivil是由golang开发的全新web框架，灵感来源于Laravel及Symfony。
- [ecgo](https://github.com/tim1020/ecgo)ecgo是一个易学、易用、易扩展的goweb开发框架
- [Gin](https://github.com/gin-gonic/gin)Gin是一个用Go语言开发的Web框架，提供类Martini的API，但是性能更好。因为有了httprouter性能提升了40倍之多。
- [Melody](https://github.com/olahol/melody)Melody是一个Go语言的微型WebSocket框架，基于github.com/gorilla/websocket开发，
- [utron](https://github.com/gernest/utron)utron是一个Go语言轻量级的MVC框架，用于快速构建可伸缩以及可靠的数据库驱动的Web应用。
- [Lessgo](https://github.com/lessgo/lessgo)Lessgo是一款Go语言编写的简单、稳定、高效、灵活的web完全开发框架。它的项目组织形式经过精心设计，实现前后端分离、系统与业务分离，完美兼容MVC与MVVC等多种开发模式，非常利于企业级应用与API接口的开发。当然，最值得关注的是它突破性地支持了运行时路由重建，开发者可在Admin后台轻松实现启用/禁用模块与操作，添加/移除中间件等功能！同时，它推荐以HandlerFunc与MiddlewareFunc为基础的函数式编程，也令开发变得更加灵活富有趣味性。
- [Hopen](https://github.com/who246/hopen)Golangweb极速开发框架。
- [Faygo](https://github.com/henrylee2cn/faygo)Faygo是一款快速、简洁的GoWeb框架，可用极少的代码开发出高性能的Web应用程序（尤其是API接口）。只需定义structHandler，Faygo就能自动绑定、验证请求参数并生成在线API文档。
- [beego](https://github.com/astaxie/beego)beego是一个用Go开发的应用框架，思路来自于tornado，路由设计来源于sinatra，[beego入门文档](https://my.oschina.net/astaxie/blog/124040)[Beego源码分析](https://my.oschina.net/fuxiaohei/blog/228999)
- [Revel](https://github.com/revel/revel)Revel是一个高生产力的Go语言Web框架。Revel框架支持热编译，当编辑、保存和刷新源码时，Revel会自动编译代码和模板；全栈特性，支持路由、参数解析、缓存、测试、国际化等功能。[Revel中文社区](http://www.gorevel.cn)[一步一步学习RevelWeb开源框架](http://www.cnblogs.com/ztiandan/archive/2013/01/17/2864498.html)
- [Martini](https://github.com/codegangsta/martini)Martini是一个非常新的Go语言的Web框架，使用Go的net/http接口开发，类似Sinatra或者Flask之类的框架，你可使用自己的DB层、会话管理和模板。
- [Tango](https://github.com/lunny/tango)Tango，微内核可扩展的Go语言Web框架。同时支持函数和结构体作为执行体，插件丰富。[Tango，微内核可扩展的Go语言Web框架](http://www.golangtc.com/t/54ae8b60421aa9396a000204)
- [Macaron](https://github.com/go-macaron/macaron)Macaron是一个具有高生产力和模块化设计的GoWeb框架。框架秉承了Martini的基本思想，并在此基础上做出高级扩展。
- [Web.go](https://github.com/hoisie/web)web.go跟web.py类似，但使用的是Go编程语言实现的Web应用开发框架。Go发布没多久该框架就诞生了，差不多是最早的Go框架。目前已经有段时间没有更新了。不过，该框架代码不多，其源码可以读一读。
- [Echo](https://github.com/labstack/echo)Echo是个快速的HTTP路由器（零动态内存分配），也是Go的微型Web框架。
- [web](https://github.com/hoisie/web)
- [goku](https://github.com/QLeelulu/goku)[goku官网](http://qleelulu.github.io/goku)

## 区块链技术

- [fabric](https://github.com/hyperledger/fabric)Fabric是一个开源区块链实现，开发环境建立在VirtualBox虚拟机上，部署环境可以自建网络，也可以直接部署在BlueMix上，部署方式可传统可docker化，共识达成算法插件化，支持用Go和JavaScript开发智能合约，尤以企业级的安全机制和membership机制为特色。你要是不知道这些术语什么意思，就记住一点，Fabric之于区块链，很可能正如Hadoop之于大数据。
- [go-ethereum](https://github.com/ethereum/go-ethereum)go-ethereum客户端通常被称为geth，它是个命令行界面，执行在Go上实现的完整以太坊节点。通过安装和运行geth，可以参与到以太坊前台实时网络并进行以下操作：a.挖掘真的以太币b.在不同地址间转移资金c.创建合约，发送交易d.探索区块历史e.很多其他功能
- [chain](https://github.com/chain/chain)**金融领域的区块链项目**.Chain是由一家刚成立两年的美国创业公司Chain推出，是一个企业级的区块链平台架构，可以让机构构造从零开始更好的金融服务。Chain开放标准在以下方面实现突破：•全新的共识模型在不到一秒的时间里实现交易的最终完成，即便是交易量非常大也能支持;•私密解决方案对区块链数据进行加密，并让相关对手方和监管者进行有选择的读取;•智能合约框架和虚拟机支持简单的规则执行，以及进行键值存储的图灵完整程序;•可伸缩的数据模型可以为网络参与者降低运行负荷;•丰富的元数据层可支持满足KYC（了解你的客户）和AML（反洗钱）要求

## 其它

- [kone](https://github.com/xjdrew/kone)可用于家庭或者企业网络的透明代理，可用来翻墙等
- [KodeRunr](https://github.com/jaxi/koderunr)KodeRunr(读作coderunner)是款我在闲暇时间用Go语言编写的应用。顾名思义，你可以用它在网页上、命令行里写程序，贴代码，与此同时无需在本地安装任何编程语言。支持Ruby,Python,GO,Swift,C,Elixir等
- [godaemon](https://github.com/tim1020/godaemon)godaemon是用来为应用增加daemon和graceful的。
- [Gomobile](https://github.com/golang/mobile)Gomobile是一个应用于iOS和Android的优秀跨平台开发库，为开发者提供用于创建Android和iOS移动平台代码的工具。
- [gojieba](https://github.com/yanyiwu/gojieba)"结巴"中文分词的Golang语言版本。
- [Cherry](https://github.com/rafael-santiago/cherry)Cherry是一个使用Go语言开发的Web聊天引擎。
- [MailSlurper](https://github.com/mailslurper/mailslurper)MailSlurper是一个便携的SMTP邮件服务器，对本地和团队应用开发来说非常有用。MailSlurper体积小运行快速，支持SQLite,MSSQL和MySQL.数据库。
- [RobustIRC](https://github.com/robustirc/robustirc)RobustIRC是不会有网络中断情况的IRC。RobustIRC主要特性：服务器不可用的时候不会有网络中断；可以使用标准IRC客户端；健壮，可以很好处理客户端和网络的连接问题
- [Qor](https://github.com/qor/qor)Qor是基于Golang开发的的CMS一系列工具库，基于Qor可以快速搭建网站的后台管理系统。Qor的工作库包含：1，后台管理：可以对数据库进去CURD管理，支持一对一，一对多，多对多关联关系维护等等；2，支持上传图片到云以及filesystem，resize、crop图片等等；3，Publish发布系统，可以修改数据，并且经过检查后，再发布到正式环境中；4，状态机，可以用于开发工作流的系统；5，I18n，翻译，可以通过在WEB界面翻译，并将翻译保存到数据库中；6，L10n，本地化，不同于翻译，他可以针对某个地区来对内容，或者数据结构进行本地化。7，Roles，权限管理；8，Exchange，通过Excel，CSV导入导出数据；9，Worker，后台任务管理，可用于跑定时任务等等
- [FishChat](https://github.com/oikomi/FishChatServer)FishChat（鱼传——鱼传尺素）分布式可伸缩IM服务器，是一款纯golang编写优秀的即时通讯软件(IM),它集合了市面上已有产品的优点,并具备智能硬件网关管理(学习QQ物联思想,构思中)。
- [goRBAC](https://github.com/mikespook/gorbac)goRBAC为Go语言应用提供了轻量级的基于角色的访问控制。
- [hey](https://github.com/rakyll/hey)Boom是google一女工程师使用Go语言开发的类似apacheab的性能测试工具。相比ab，boom跨平台性更好，而且更容易安装。
- [Mattermost](https://github.com/mattermost/platform)mattermost是一个Slack的开源替代品。Mattermost采用Go语言开发，这是一个开源的团队通讯服务。为团队带来跨PC和移动设备的消息、文件分享，提供归档和搜索功能。
- [glot](https://github.com/prasmussen/glot)glot是可以可以在线运行各种编程语言代码片段的平台，项目采用HaskellScript、Go、Erlang和Shell开发，运行环境基于Docker容器进行。
- [Lantern](https://github.com/getlantern/lantern)Lantern是一个点对点科学上网软件。
- [dog-tunnel](https://github.com/vzex/dog-tunnel)狗洞是一个高速的P2P端口映射工具，同时支持Socks5代理。0.5版后开始开源，UDP底层基于开源库KCP重写，效率大大提高，在恶劣环境下优势明显。同时提供非P2P版本（Lite版本），两端连接过程完全不依赖中间服务器，支持加密和登陆认证，自动重连，但是需要人为确保两端能正常连通（否则请使用默认的P2P版本）
- [GRPC](https://github.com/grpc)GRPC是一个高性能、开源和通用的RPC框架，面向移动和HTTP/2设计。目前提供C、Java和Go语言版本，分别是：grpc,grpc-java,grpc-go.其中C版本支持C,C++,Node.js,Python,Ruby,Objective-C,PHP和C#支持.GRPC基于HTTP/2标准设计，带来诸如双向流、流控、头部压缩、单TCP连接上的多复用请求等特。这些特性使得其在移动设备上表现更好，更省电和节省空间占用。
- [LiteIDE](https://github.com/visualfc/liteide)LiteIDE是一款开源、跨平台的轻量级Go语言集成开发环境（IDE）。
- [firefly-proxy](https://github.com/yinghuocho/firefly-proxy)穿墙工具。GFW梯子。提供客户端和服务端。支持多个平台，包括linux，macos，windows，android
- [wu](https://github.com/shanzi/wu/)一个监听文件变化并自动执行某些操作的小工具，可以用于配置修改后自动重启webserver
- [apex](https://github.com/apex/apex)管理，部署awslambda函数的工具，支持用go语言编写lambda函数（注：目前aws官方不支持用go语言编写lambda函数，但是apex却可以变相支持）
- [gosuv](https://github.com/codeskyblue/gosuv)进程管理，类似于python的supervisord，提供了web管理界面
- [chaosmonkey](https://github.com/Netflix/chaosmonkey)ChaosMonkey是netflix公司开源的一个用于服务可用性测试的工具，通过有计划的在生产系统制造真实的故障（如cpu负载高，内存溢出，磁盘写满，服务器宕机等）来检测系统的可用性。
- [scheduler](https://github.com/shotdog/scheduler)scheduler专门进行任务的调度分发任务工作，各个任务的具体任务执行分配到各个项目中，从而达到对任务的统一配置和管理。该工具提供了web管理界面
- [hugo](https://github.com/gohugoio/hugo)Hugo是由Go语言实现的静态网站生成器；简单、易用、高效、易扩展、快速部署；相比于Hexo、Jekyll，hugo的优势是生成速度极快。
- [Lime](https://github.com/limetext/lime)相对上面的几款Go语言在云端和服务器端之外，Lime则显得比较特殊。Lime，则是一款用Go语言写的桌面编辑器程序，被看做是著名编辑器SublimeText的开源实现。
- [Gogs](https://github.com/gogits/gogs)Gogs则是一款由国人无闻（GitHub）开发的自助Git服务项目。Gogs的目标是打造一个最简单、最快速和最轻松的方式搭建自助Git服务。据作者称，之所以选择使用Go语言开发，就是Go允许Gogs可以通过独立的二进制分发，且对跨平台有良好支持。
- [蚂蚁笔记](https://github.com/leanote)
- [MinDoc](https://github.com/lifei6671/mindoc)

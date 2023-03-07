# DDD 项目（六边形架构）

## 层级说明
```base
├── interfaces 接入层 【http/grpc适配, 这一层调用application层; 比如interfaces层定义了输入层的相关方法，以使用gin提供http接口为例，这里的handler等为使用gin提供的一些http接口，这一层调用application层】
│   ├── grpc
│   └── http
│   └── facade  引用其他微服务（接口防腐层）
│   ├── event 事件
│   │   └── subscribe mq消费入口
│   ├── job 定时任务
├── application 应用层 【主要是调用domain层与infrastructure层来实现功能】
│   ├── assembler   负责将内部领域模型转化为可对外的DTO
│   └── dto Application层的所有接口返回值为DTO -- 入参/出参
│   └── service 负责业务流程的编排，但本身不负责任何业务逻辑
├── domain 领域层 【主要是定义了entity，以及repository接口；entity里头会包含一些领域逻辑,Domain模块仅依赖Types模块】
│   ├── aggregate 聚合 【对于需要两个repo一起操作的，可以进行聚合，比如创建用户的时候有userRepo,还有日志的userLogRepo】
│   ├── entity 实体 业务逻辑。也可以参数校验，扩展一些简单方法，减轻service的压力
│   ├── event 事件
│   │   ├── publish 所有发送mq在此处理
│   │   └── subscribe 所有接受到mq处理逻辑在此处理
│   ├── irepository 接口
│   ├── service 领域服务 【单一操作，比如查看用户信息。没有聚合的操作的时候，在此实现】
└── infrastructure 基础设施层 【这里提供了针对domain层的repository接口的实现，还有其他一些基础的组件，提供给application层或者interfaces层使用】
│   ├── config 配置文件
│   ├── consts 系统常量
│   ├── pkg 常用工具类封装（DB,log,util等）
│   └── repository 针对domain层的repository接口的实现
│   │   └── converter domain内对象转化 po {互转}
│   │   └── dao 针对domain层的repository接口的具体实现
│   │   └── po 数据库映射对象
└── types 完全独立的模块(DP)，封装自定义的参数类型,例如 phone 相关的类型，校验合法、区号等。  

```

## DDD 小结
DDD 一般分为 interfaces、application、domain、infrastructure 这几层；其中 domain 层不依赖其他层，它定义 repository 接口，infrastructure 层会实现；application 层会调用 domain、infrastructure 层；interfaces 层一般调用 application 层或者 infrastructure 层。



## 相关概念

- DDD 等相关概念：https://domain-driven-design.org/zh/ddd-concept-reference.html

- VO（View Object）：视图对象，用于展示层，它的作用是把某个指定页面（或组件）的所有数据封装起来。
- DTO（Data Transfer Object）：数据传输对象，这个概念来源于 J2EE 的设计模式，原来的目的是为了 EJB 的分布式应用提供粗粒度的数据实体，以减少分布式调用的次数，从而提高分布式调用的性能和降低网络负载，但在这里，我泛指用于展示层与服务层之间的数据传输对象。
- DO（Domain Object）：领域对象 (entity)，就是从现实世界中抽象出来的有形或无形的业务实体。
- PO（Persistent Object）：持久化对象，它跟持久层（通常是关系型数据库）的数据结构形成一一对应的映射关系，如果持久层是关系型数据库，那么，数据表中的每个字段（或若干个）就对应 PO 的一个（或若干个）属性。

> ```
> 用户发出请求（可能是填写表单），表单的数据在展示层被匹配为 VO。
> 展示层把 VO 转换为服务层对应方法所要求的 DTO，传送给服务层。
> 服务层首先根据 DTO 的数据构造（或重建）一个 DO，调用 DO 的业务方法完成具体业务。
> 服务层把 DO 转换为持久层对应的 PO（可以使用 ORM 工具，也可以不用），调用持久层的持久化方法，把 PO 传递给它，完成持久化操作。
> 对于一个逆向操作，如读取数据，也是用类似的方式转换和传递
> ```



# 项目运行

## 配置

manifest/config/config.yaml

```yaml
# HTTP Server.
server:
  address: ":8018"
  serverName: "garden-be"
  mode: "release"

# Database.
mysql:
  username: "root"
  password: "root"
  dbHost: "10.4.7.71"
  dbPort: 3307
  dbName: "test"

# Log.
log:
  logFileDir: "./tmp"
  appName: "garden-be"
  maxSize: 512
  maxBackups: 64
  maxAge: 7

```

infrastructure/config/config.go 注意文件路径修改

## 运行

- 直接运行：go run main.go
- 编译：go build

tips:
通过 DDD 可以对中台和微服务起到承上启下作用，中台是更偏业务角度考虑，DDD 通过领域建模拆解、划分业务领域边界，指导微服务的落地。在使用 DDD 过程中有一些注意点也需要关注下：

DDD 适合偏复杂业务，DDD 不是万能的。简单业务使用 DDD 会有些杀鸡用牛刀感觉（思考架构三原则：简单、合适、演进），不要拿着 DDD 这个锤子到处找钉子；
DDD 分层建议采用严格分层，不跨层调用，而是采用依赖注入方式把相关实例传入下层（例如不要从接口层直接调用存储层方法，因为跨层调用会导致整个调用链变复杂）；
DDD 目录结构命名，这块也是比较关键一点。目前 Go 是倾向简洁，不希望向 Java 那么冗余，所以这块命名还可以在 DEMO 基础上进一步优化；
DDD 分层会接口一多，代码可读性不好的问题。可以通过好的命名来规避（比如统一后缀、选取合适简短的接口名），同时用依赖倒置思维逐层看接口，以及其依赖；
DDD 设计步骤，可以按领域层 -> 基础层 -> 应用层 -> 接口层，一般是按这个步骤开发；
DDD 分层后，每层隔离得比较干净，非常适合单元测试和 Mock 测试（可以参考文末 food-app-server 这个仓库）

reference: https://github.com/jettjia/go-ddd

STYLEGUIDE.md 部分参考：
https://zhuanlan.zhihu.com/p/91525839
https://tech.meituan.com/2017/12/22/ddd-in-practice.html
https://tkstorm.com/posts-list/cloud-native/ddd-layer/

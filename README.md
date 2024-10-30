# 软件介绍:
定海WAF网站防火墙适用于个人开发者,工作室,小公司的开源防火墙.保护网站安全,降低网站漏洞带来的数据泄露等风险.致力于将定海开发为一个更加灵活,性能高效,更加安全的防火墙.

**初衷:**

    1. **便捷: **如果是单纯的从自己的日志或者nginx,apache日志等查看信息不方便 , 不知道自己的网站到底谁在访问,请求了什么. waf在网站或者API防护可以更方便用户查询这些信息,做出及时的处理.
    2. **DIY: **在开发过程中一些特定的功能加入自己的想法.

## 技术栈
项目上层服务使用kratos框架进行开发,内核服务使用coraza waf开源引擎,提高内核开发效率.

> 需要掌握的知识:
>

1. golang [【置顶】Go语言学习之路/Go语言教程 | 李文周的博客](https://www.liwenzhou.com/posts/Go/golang-menu/)
2. kratos  [简介 | Kratos](https://go-kratos.dev/docs/)
3. resetful [重新认识RESTful | 少个分号](https://shaogefenhao.com/column/restful-api/restful-api-introduction.html#restful)
4. grpc [Basics tutorial | Go | gRPC](https://grpc.io/docs/languages/go/basics/)
5. proto3+protobufvalidate [Language Guide (proto 3) | Protocol Buffers Documentation](https://protobuf.dev/programming-guides/proto3/)
6. coraza waf [OWASP Coraza - Enterprise-grade open source web application firewall library](https://coraza.io/)
7. mysql + gorm [GORM 指南 | GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://gorm.io/zh_CN/docs/index.html)
8. redis
9. kafka
10. etcd
11. clickhouse [什么是ClickHouse？ | ClickHouse Docs](https://clickhouse.com/docs/zh)

# 项目结构
## 整体结构
![](https://cdn.nlark.com/yuque/0/2024/png/34606362/1730189896389-59aabf81-8acf-4017-ab97-454ed92a4450.png)

尽可能的保证waf内核的轻量级,避免出现过多的冗余功能,保证waf内核的稳定性. 用户通过上层服务去设置防护配置,利用etcd的watch属性进而修改内置中waf实列的配置.

## 上层服务结构
![](https://cdn.nlark.com/yuque/0/2024/png/34606362/1730202320311-dd969cda-fe2b-4242-9544-4c60ef32d4ff.png)

### 目录结构
```plain
├─api       #/ 下面维护了微服务使用的proto文件以及根据它们所生成的go文件
│  ├─dashBorad # 数据看板
│  │  └─v1
│  ├─user  # 用户相关服务
│  │  └─v1
│  └─wafTop # 防护配置
│      └─v1
├─app  # app
│  ├─dashBorad # 数据看板相关服务
│  │  ├─cmd # // 整个项目启动的入口文件
│  │  │  └─dashBorad
│  │  ├─configs
│  │  ├─internal // 该服务所有不对外暴露的代码，通常的业务逻辑都在这下面，使用internal避免错误引用
│  │  │  ├─biz // 业务逻辑的组装层
│  │  │  ├─conf // 内部使用的config的结构定义，使用proto格式生成
│  │  │  ├─data // 业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口。
│  │  │  ├─server // http和grpc实例的创建和配置
│  │  │  └─service // 实现了 api 定义的服务层，类似 DDD 的 application 层
│  │  └─third_party
│  │      ├─errors
│  │      ├─google
│  │      │  ├─api
│  │      │  └─protobuf
│  │      │      └─compiler
│  │      ├─openapi
│  │      │  └─v3
│  │      └─validate
│  ├─user # 用户相关服务
│  │  ├─cmd // 整个项目启动的入口文件
│  │  │  └─user
│  │  ├─configs
│  │  ├─internal
│  │  │  ├─biz
│  │  │  │  └─iface
│  │  │  ├─conf
│  │  │  ├─data
│  │  │  │  └─model
│  │  │  ├─server
│  │  │  │  └─plugin
│  │  │  └─service
│  │  └─third_party
│  │      ├─errors
│  │      ├─google
│  │      │  ├─api
│  │      │  └─protobuf
│  │      │      └─compiler
│  │      ├─openapi
│  │      │  └─v3
│  │      └─validate
│  └─wafTop # 上层防护配置相关服务
│      ├─cmd
│      │  └─wafTop
│      ├─configs
│      └─internal
│          ├─biz # 业务逻辑
│          │  ├─iface
│          │  ├─rule
│          │  ├─site
│          │  └─strategy
│          ├─conf
│          ├─data # 数据处理
│          │  ├─dto
│          │  └─model
│          ├─server
│          │  └─plugin
│          └─service # 
│              ├─rule
│              ├─site
│              └─strategy
└─third_party  # 依赖的外部porto文件
    ├─buf
    │  └─validate
    │      └─priv
    ├─errors
    ├─google
    │  ├─api
    │  └─protobuf
    │      └─compiler
    ├─openapi
    │  └─v3
    └─validate

```

### 数据库ER图
![](https://cdn.nlark.com/yuque/0/2024/png/34606362/1730252770131-13ba0a2d-f4dc-4706-a1a3-45c75cf47122.png)

# 项目更新日志
> v1.0.0
>

1. 内核具有基础防护功能 , 支持动态更新策略 , 以及各个网站之间的waf实列隔离,简单的实现了反向代理,将请求转发至真正的后端地址支持coraza内置防护策略,<font style="color:rgb(31, 35, 40);">包括：SQL注入（SQLi），跨站点脚本（XSS），PHP和Java代码注入，HTTPoxy，Shellshock，脚本/扫描器/机器人检测&元数据和错误泄漏。同时也支持用户自定义防护策略,v1.0.0版本实现了用户自定义IP防护策略.可以根据不同策略的定义,对恶意请求作出是否拦截的处理.可以将攻击事件落地本地磁盘的csv文件.</font>
2. <font style="color:rgb(31, 35, 40);">上层服务,用户配置需要防护的网站,策略,规则组以及自定义规则. 目前内置规则不支持用户自定义.</font>

> 部署说明:
>

1 .  MySQL: 安装mysql8版本以上

<font style="color:rgb(31, 35, 40);"></font>


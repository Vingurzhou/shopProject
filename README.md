# 仿商城秒杀系统(Golang练手项目)


---

这个项目是前后端一起写，返回的大多数`Html`（当然接口也是可以的）

主要分为了`前台`，`后台`，`秒杀`这个**三个模块**，大多数比较基础但是还是能学到思想

下面简单说一下**技术栈**，**架构图**(重新手画)，**目录**，**学习经验**，**如何使用**。

##  技术栈/架构图

> 其实里面的分布式验证可以用Redis实现 但是只要 一致性哈希 这个功能完全可以用代码层面实现
>
> SLB 业务不大的情况下 完全可以自己配置个组 但是在一个内网做均衡会快很多

|        技术         |                           相关文档                           |
| :-----------------: | :----------------------------------------------------------: |
| 消息队列： RabbitMQ | [中文文档](https://blog.csdn.net/weixin_51485807/article/details/122761910) |
|    MVC框架：Iris    |          [中文文档](https://www.topgoer.com/Iris/)           |
|   静态加速： CDN    | [腾讯云CDN的详细配置](https://cloud.tencent.com/developer/article/1462593?from=15425) |
|    负载均衡：SLB    |       [阿里云SLB](https://www.aliyun.com/product/slb)        |
|    数据库：Mysql    | [中文文档](https://www.docs4dev.com/docs/zh/mysql/5.7/reference/) |
|  数据库映射：GORM   |    [中文文档](https://gorm.io/zh_CN/docs/index.html)     |
|    压力测试：Wrk    | [中文文档](https://segmentfault.com/a/1190000023212126) |

![](https://raw.githubusercontent.com/HengY1Sky/GoSecKillMall/main/GoFramework.webp)

##  教学目录/框架目录

*教学地址：https://coding.imooc.com/class/chapter/347.html

```
├── backend
│   ├── main.go # 后台启动文件
│   └── web
├── common # 公共库
├── fronted
│   ├── main.go # 前端启动
│   ├── middleware # 前端中间件
│   ├── productMain.go # CDN入口
│   └── web # 逻辑
├── rabbitmq # rabbitmq
├── datamodels # 基础模型
├── repositories # 面向SQL
├── services # 服务
├── consumer.go # 抢购队列消费
├── getOne.go # 判断是否超卖
├── go.mod
├── go.sum
└── validate.go
```


## 经验/更新


---

1. 使用的是`iris/v12`最新的框架，并且使用了包管理，开箱即用。
2. 将原生的SQL写法更换成了Gorm，更加规范了。
3. 完善抢购逻辑，将整个抢购思路形成闭环。

## 快速开始


- SQL文件直接到`Navicat`**创建自己的库**然后**创建查询**然后**执行**就好了

```bash
$ export GO111MODULE=on && export GOPROXY=https://goproxy.cn # 开启mod模块以及换源
$ go mod tidy
# 进入Goland开始操作吧
```


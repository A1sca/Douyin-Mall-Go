# Douyin-Mall-Go

## 项目初始化
按照官方文档的idl, 对项目进行初始化


# 服务模块实现
> 所有配置都只改动了test的配置, 没有改动dev和online的配置

## 用户服务模块
根据教程写一次基本的登陆注册

步骤为: 
1. 确定需要提供什么服务, 写idl
2. 用cwgo生成服务的基本代码
3. 写数据模型(位于biz/model)
4. 可以在dal/mysql里面, 用gorm的auto migrate创建表
4. 写业务逻辑(位于biz/service)

> 写完代码之后, 可以执行测试, 但是测试的时候如果要加载配置文件, 会出现路径错误, 找不到配置文件. 不好解决, 建议不要运行那几个测试

运行一个服务需要注意的东西:
1. 在main.go的kitexInit中, 向consul注册服务. 
2. consul在哪个端口? 这个根据自己的docker compose看.
3. 服务运行在哪个端口? conf下面的yaml文件里配置
4. main.go默认不执行dal的init, 需要自己添加

## 订单服务模块
运行在端口8084
1. 修改conf的mysql dsn
2. main.go里面添加dal.Init(), mysql的Init()中, augomigrate新增表
3. main.go的kitexInit()中, 向注册中心注册服务
4. 指定服务的名字为 order, 需要修改配置文件. main.go的模版代码已经读取了这个配置

```yaml
kitex:
  service: "order"
```

```go
opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
    ServiceName: conf.GetConf().Kitex.Service,
}))
```


## api
运行在端口8080
1. 运行makefile, 生成模版代码
2. conf的yaml文件添加consul的地址, conf.go的Hertz配置添加一个新成员: 注册中心地址
3. 添加一个rpc包, 用于初始化各个rpc服务
4. 初始化rpc服务的时候, 需要指定rpc服务的名字, 怎么知道这个名字?? (服务端的配置文件可以指定服务名字)

# Douyin-Mall-Go

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

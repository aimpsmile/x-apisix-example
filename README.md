## x-apisix-example简介
* x-apisix的教程工程
* 希望go-micro使用用户可以快速入门
* [项目路径](https://github.com/micro-in-cn/x-apisix)

### 环境变量兼容go-micro(请修改Makefile)
* MICRO_REGISTRY=etcd
* MICRO_REGISTRY_ADDRESS=etcd.service:2379

## 版本匹配逻辑
> 网关模板匹配规则（范围不允许交叉重叠）
- 版本一般v1.3.2.2.2.2.2 会转成float64位进行比对。比对来讲更简单，但是float64小数位超过15位精度就丢失，请规范的使用版本
-   `>` `>=` `<` `<=` `=` 浮点数数学比对工具
-  `~ v1.3`  匹配 [v1.3.00 ,v1.33.9999999] 闭区间
- `v1.3,v1.5` 匹配 [v1.3.00,v1.5.9999999] 闭区间
- `*` 如果版本没有匹配上，保底的。
## conf.yml配置详情
```yaml
global:
  log:   # 日志
    level: debug   # 日志级别
    logFileDir: /logs  # 日志目录
    maxSize: 10  # 每个日志文件保存的最大尺寸 单位：M
    maxBackups: 3 # 日志文件最多保存多少个备份
    maxAge: 10   #文件最多保存多少天
    skip: 0
conf: # x-apisix配置
  check:   # 同步检查配置
    retres: 3 #请求接口重试次数
    interval: 5 #检查间隔：单位s
  leader:  # leader主从配置
    id: "monitor-default"  # 选举需要用的id
    group: "gateway"  #leader 属组
    nodes:  #节点列表
      - "etcd.service:2379"
  gateway:  # 网关的相关的配置
    timeout: 5000  #请求接口超时时间毫秒
    retries: 1 #请求接口重试次数
    apikey: "edd1c9f034335f136f87ad84b625c8f1"  #网关身份认证token(apisix v1.2之后有用)
    baseurl: "http://apisix.service:8888/apisix/admin" #网关的接口路径
    protoPath: "/storage/code/examples/micro-to-apisix/grpc/proto"  #proto文件路径
    forbidRoutes: # 禁用同步到网关的路由
      #	服务的状态，必须路径
      /stats: true
      #	健康检查，必须路径
      /health: true
  filter:  #同步过滤配置
    -
      bu: "aimgo"  # bu名称
      stype: "web"  # 服务类型，支持http2、web、srv、tool、api、job、apigw、webgw
      module: "*"  # *代码所有的文档
      ver:  # 版本
        "*":  #所有版本
          routeTpl: "yaml/web.routes.yaml"  #路由模板、相对配置目录即可
          upstreamTpl: "yaml/web.upstreams.yaml" #负载均衡模板、相对配置目录即可
          tplFormat: "yaml"  # 模板格式：json、yaml
          hosts:  # 解析的域名
            - "web.uqudu.com"
            - "web2.uqudu.com"
    -
      bu: "aimgo"
      stype: "http2"
      module: "*"
      ver:
        "*":
          routeTpl: "yaml/grpc.routes.yaml"
          upstreamTpl: "yaml/grpc.upstreams.yaml"
          tplFormat: "yaml"
          hosts:
            - "srv.uqudu.com"
            - "http2.uqudu.com"


```

## 执行脚本之前请修改对应的配置
* 修改 **x-apisix-example\apisix\config\local\apisix\conf.yml** 配置项目
	* gateway.baseurl # apisix的admin api路径
	* gateway.protoPath # proto路径
* host修改
	* apisix服务器ip web.uqudu.com srv.uqudu.com http2.uqudu.com web2.uqudu.com
### 启动x-apisix

```sh
	# 安装
	make xapisix-install
	# 执行
	make xapisix-exec
```

### websocket

```sh
	# 安装
	make websocket-install
	# 执行
	make websocket-exec
```
### grpc

```sh
	# 安装
	make grpc-install
	# 执行
	make grpc-exec
```
### web

```sh
	# 安装
	make web-install
	# 执行
	make web-exec
```


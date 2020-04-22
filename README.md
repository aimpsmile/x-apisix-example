## x-apisix-example简介
* x-apisix的教程工程
* 希望go-micro使用用户可以快速入门
* [项目路径](https://github.com/micro-in-cn/x-apisix)

### 环境变量兼容go-micro(请修改Makefile)
* MICRO_REGISTRY=etcd
* MICRO_REGISTRY_ADDRESS=etcd.service:2379

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



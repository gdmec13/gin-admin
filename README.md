# gin-simple-cloud-storage
> gin-simple-cloud-storage 这是一个简单的本地云存储项目，熟悉下写法。


### 使用
```shell
go mod tidy

go run cmd/main.go
```


### 目录结构
```shell
├─app
│  ├─api
│  │  └─request
│  ├─config
│  ├─global
│  ├─init
│  ├─model
│  └─router
├─cmd
├─log
├─pkg
│  ├─log
│  ├─middleware
│  ├─response
│  └─util
└─test
```
**结构目录说明**
1. `app`：包含了项目的主要应用程序代码。
   + `api`：包含API相关的代码。 
   + `request`：包含处理请求的相关代码。 
   + `config`：包含项目的配置文件。 
   + `global`：包含全局变量和常量。 
   + `init`：包含项目的初始化逻辑。 
   + `model`：包含数据模型相关的代码。 
   + `router`：包含路由相关的代码。
2. `cmd/`: 这个目录包含了你的应用程序的入口点。通常在其中包含main.go。
3. `pkg/`: 这个目录包含了你希望被其他项目引用的包。在这里，你可以将一些通用的工具函数、中间件等组织起来。
   + `log`：包含日志处理相关的代码。 
   + `middleware`：包含中间件相关的代码。
   + `response`：包含处理响应的相关代码。 
   + `util`：包含通用的工具函数。
4. `log`：包含日志文件。
5. `test`：包含测试相关的代码。

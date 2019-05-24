## 如果 go get 不动，请自行加梯子
```shell
export GOPROXY=https://goproxy.io
# 注意该梯子不支持私有repo，私有repo请去掉GOPROXY
export GOPROXY=
```

## 依赖
go >= 1.11(由于使用了 go mod 管理版本依赖)

```shell
# 如果想在GOPATH下用mod, 请设置 GO111MODULE=on 则在 GOPATH/src 目录下使用 go get 时也默认采用 go mod
export GO111MODULE=on
```

## 开始
```
// 1. 获取 yago
go get github.com/hulklab/yago

// 2. 用 yago 在当前目录创建你的项目 myapp
yago init -a myapp

// 3. 进入目录初始化
cd myapp/
go mod init
```

## 目录结构
```
.
├── cmd
│   └── main.go
├── conf
│   └── app.toml
├── main.go
├── modules
│   └── home
│       ├── homecmd
│       │   └── home.go
│       ├── homedao
│       │   └── home.go
│       ├── homehttp
│       │   └── home.go
│       ├── homemodel
│       │   └── home.go
│       ├── homerpc
│       │   ├── home.go
│       │   ├── home_test.go
│       │   ├── protobuf
│       │   │   └── homepb
│       │   │       ├── home.pb.go
│       │   │       └── home.proto
│       │   └── README.md
│       └── hometask
│           └── home.go
├── routes
│   ├── cmdroute
│   │   └── cmd.go
│   ├── httproute
│   │   └── http.go
│   ├── rpcroute
│   │   └── rpc.go
│   └── taskroute
│       └── task.go
├── third
│   └── homeapi
│       ├── home.go
│       └── protobuf
│           └── homepb
│               ├── home.pb.go
│               └── home.proto
└── tools
    └── build.sh
```

## 路由

#### 1. http 路由 
```
@reference example/app/modules/home/homehttp/home.go
```

#### 2. cmd 路由
```
@reference example/app/modules/home/homecmd/home.go
```

#### 3. task 路由
```
@reference example/app/modules/home/hometask/home.go
```

#### 4. rpc 路由
```
@reference example/app/modules/home/homerpc/home.go
```

## 配置
1. 位置: `conf/app.toml`
2. 解析: `conf.go`
3. 使用: `@reference libs/orm/orm.go line 29`

## 组件
1. 全局容器: `com.go`
2. 使用: `@reference libs/rds/redis_test.go` 

## 模块

##### 1. 模块基础目录 
```
dao model http rpc task cmd
```
##### 2. 创建新模块
使用yago创建模块会自动将路由加载到myapp/routes 下
```
cd myapp
yago new -a myapp -m newmodule
```

## 错误
```
# 系统级错误定义处
error.go
# 使用
@reference example/app/modules/homehttp/home.go::AddAction
```

## Third
1. 目录规范 `@see example/app/third`
2. http-api 使用样例 `@reference example/app/third/homeapi/home.go`

## Goland 使用 mod

1. Preferences -> Go -> Go modules(vgo)
2. ![](http://p406.qhimgs4.com/t0100eba6c9f82cb921.png)
3. 如果还有标红的提示，点击 Sync packages
![](http://p406.qhimgs4.com/t019f0fcae328f7a0e0.png)

# go-sample

## 问题

1. supervisor 脚本如何优雅的管理
    - 如果supervisor任务不打二进制包,直接运行代码,留存superviosr配置文件即可, 更新supervisor只需要拉代码,重启脚本即可
    - 如果supervisor任务打成二进制包,就麻烦了,如何判断二进制包基于哪个git版本呢

2. 代码版本管理,以及二进制文件如何查看是由哪个版本产生的

## 参考文献

1. 项目文件夹命名 <https://github.com/golang-standards/project-layout/blob/master/README_zh-CN.md>
2. 项目结构,参考了kratos新版文档 <https://go-kratos.dev/docs/intro/layout> 
3. 异常处理,参考了毛剑的教程,kratos@v1.0 kratos@v2.0 <https://go-kratos.dev/docs/component/errors>
4. 

## 使用的工具
1. goimports `go install golang.org/x/tools/cmd/goimports@v0.14.0`
2. golangci-lint `curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.0`
3. gmchart `go install github.com/PaulXu-cn/go-mod-graph-chart@v0.5.3`
4. 


## 使用的第三方库

1. gin
2. grpc
3. grpc-gateway   https://github.com/grpc-ecosystem/grpc-gateway

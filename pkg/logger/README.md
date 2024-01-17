# logger

## logger 选择
最初考虑使用go 官方库 (1.20后) log/slog
能够使用json 打印出struct
不过这个日志库,关于打印error stacktrace 貌似弱一些

还是选择了 zap


## 日志切分
基本上大部分的logger库,都支持输入一个 io.Writer 来存储日志
而日志切分的一些库,都实现并提供一个io.Writer
还是选择了一贯的lumberjack

至于日志切分的责任划分,和运维讨论了一下,还是尽量由开发来进行切分,总体日志不要超过1g

目前新上的所有代码,基本都是在k8s中启动
如果没有特殊的需求,k8s的pod 都没有外挂存储,所以所有的内容都是在k8s的pod内存中,重启就会丢失内存中的所有数据
所以不论是 我们的日志落盘,还是 stdout 输出到>>xxx.log,本质都是落入到pod内存中
而stdout 输出到>>xxx.log 没有切分功能,故放弃
而直接输出stdout,不落文件,虽然也可以上传到ali云,但是如果上传失败,就没有办法再找到日志了,也放弃
如果有永久化需求,需要申请一块云存储硬盘,挂载到pod上,日志落盘到云存储硬盘,才可以

## 与context的结合

海外方案:
在请求到达的时候,先使用
ExtraField(),把context里面的所有字段解析成[]zap.Field
然后使用 NewContext()
使用metadata.FromContext(ctx),把context里面的metadata(存储自定义字段的玩意,一个map)取出来
clone 一个zap logger ,丢到metadata
其实这个时候每个请求里面的子logger 就已经有ctx的一些关键字段了

``` go
// With creates a child logger and adds structured context to it. Fields added
// to the child don't affect the parent, and vice versa.
```


使用的时候,使用 WithContext()方法,返回一个带这些字段的logger
当在别的模块使用这个logger的时候,就携带了traceid
不过把 child logger 直接丢到ctx,是否是最佳实践?





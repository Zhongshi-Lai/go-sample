# di

建立di这个文件夹的核心目的是为了
存放 
wire.go wire_gen.go
如果 wire.go wire_gen.go 和main.go 在一个文件夹
那么 go run main.go 会报错, undefined: InitializeApp()
找不到wire_gen.go 的内容

找不到 wire_gen.go 的内容的原因可能是
wire_gen.go 和wire.go 有下面的标签
//go:build wireinject
// +build wireinject


## 冲突问题
InitializeApp() 这个方法,出现在
wire_gen.go 和 wire.go 在一个package内
但是为什么没报出冲突呢
就是 wire.go
需要加标签

//go:build wireinject
// +build wireinject



todo: server 是否需要放在这里
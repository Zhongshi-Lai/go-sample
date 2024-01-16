# di 

编写wire.go文件的时候
去掉
//go:build wireinject
// +build wireinject

才能正常编辑,比如自动import包,检查错误

运行wire二进制文件之后,生成wire_gen.go 之后,记得加回去

wire.go wire_gen.go 如果和main在一个包内,启动 go run main.go会找不到wire_gen.go里面的函数,就单独丢一个包了
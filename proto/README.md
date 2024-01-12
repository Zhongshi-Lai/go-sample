# proto

proto 文件写法

option go_package = "sample.v2";
The import path must contain at least one period ('.') or forward slash ('/') character.

而这里的 sample.v2 变成pb.go文件之后,会变成包名
package sample_v2

--proto_path=proto 表示从proto目录下读取proto文件。
--go_out=proto 表示生成的Go代码保存的路径。
--go_opt=paths=source_relative 表示输出文件与输入文件放在相同的相对目录中。
book/price.proto 表示在proto目录下的book/price.proto文件。


```bash

protoc -I . \
    --go_out ../api_gen --go_opt paths=source_relative \
    --go-grpc_out ../api_gen --go-grpc_opt paths=source_relative \
    ./sample/v1/sample.proto

```
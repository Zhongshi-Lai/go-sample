module go-sample

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	google.golang.org/grpc v1.57.2
	google.golang.org/protobuf v1.30.0
	github.com/golang/protobuf v1.5.3
	google.golang.org/genproto v0.0.0-20230526161137-0005af68ea54 // indirect
)

replace (
	golang.org/x/mod => golang.org/x/mod v0.13.0
	golang.org/x/sys => golang.org/x/sys v0.13.0
	golang.org/x/tools => golang.org/x/tools v0.14.0
)


# err处理

## 需求:

1. 异常有堆栈,能追踪到事发现场
2. 统一打印异常日志,不要散落在业务代码中
3. 异常能够被直接返回给上层 (在没有特殊需求的情况下,调用方无需处理异常,直接return err 即可)
4. 异常能够追踪根因
5. 异常可以携带一些参数,比如trace_id
6. 异常兼容grpc框架
7. 兼容http框架


打印日志的时候,打印error 然后再打印根因(因为根因可能没有trace stack)

## 实现
总体上:参考了 kratos V2 版本的error实现
https://go-kratos.dev/docs/component/errors

### 定义结构体
```go

type Status struct {
	Code     int32  // 错误码，跟 http-status 一致，并且在 grpc 中可以转换成 grpc-status
	Reason   string   // 错误原因，定义为业务判定错误码
	Message  string // 错误信息，为用户可读的信息，可作为用户提示内容
	Metadata map[string]string // 错误元信息，为错误添加附加可扩展信息
}

// Error is a status error.
type Error struct {
    Status
    cause error
}

```

### 使用异常

不管是纯业务异常,还是数据库异常, error都是在业务代码中new出来的,而不是预先定义好的结构体
所以追踪堆栈信息的时候,是能追踪到 errors.New() 这一行
也就能找到案发现场,满足[需求1]


#### 纯业务异常
举例: 订单超过支付时限; 订单状态不对 ; 未到开售时间
这些业务异常,没有更深层次的异常了(数据库,redis等没有报错)

err.WithMetadata() 可以附著一下信息,满足[需求5]

```go

return errors.New(500, "USER_NAME_EMPTY", "user name is empty")


// 如果需要附加一下其他信息
err := errors.New(500, "USER_NAME_EMPTY", "user name is empty")
err = err.WithMetadata(map[string]string{
"foo": "bar",
})

```

#### 包含数据库/redis等中间件报错
err.WithCause() 可以包含根因,满足[需求4]


```go

order ,err := dao.getOrderByOrderNo()
if err != nil {
    newErr := errors.New(500, "USER_NAME_EMPTY", "user name is empty")
    newErr = err.WithCause(err)
	return newErr
}

```


#### error转换

一般我们写函数,接受的error 都是一个 error-interface,即实现了 Error() string 方法的一个普通error
如何转换成
type Error struct {
Status
cause error
}

这个结构体呢?

使用方法 func FromError(err error) *Error

他会首先使用 标准库 errors.As方法,根据指针来判断,到底是不是一个 type Error 类型的error
As finds the first error in err's chain that matches target

如果不是 type Error
再利用 grpc提供的 gs, ok := status.FromError(err) 
尝试判断这是否是一个grpc 类型的err
如果是grpc类型的err,也认,兼容了,转换成 type Error



##### 被grpc框架承认
同时,这个 func FromError(err error) *Error ,把普通error转换为 type Error 的操作,在grpc的框架里面也有,kratos应该是仿写的grpc

grpc也有一个方法, func FromError(err error) *Error ,只要你实现了 type grpcstatus interface{ GRPCStatus() *Status } 这个interface
就可以被转换为 status.Status ,被grpc承认
因此也就满足了 [需求6]

##### 被http框架承认

只要error 被转换为 type Error 
就包含了
```go

type Status struct {
	Code     int32  // 错误码，跟 http-status 一致，并且在 grpc 中可以转换成 grpc-status
	Reason   string   // 错误原因，定义为业务判定错误码
	Message  string // 错误信息，为用户可读的信息，可作为用户提示内容
	Metadata map[string]string // 错误元信息，为错误添加附加可扩展信息
}
```


包含了http状态码,在middleware 稍微处理一下即可返回给http框架
因此也就满足了 [需求7]



### 打印异常

1. middleware 层 对error 进行转换,转换成 type Error
2. 打印当前接受到的异常,这个异常是包含堆栈的,能够追溯到new
3. 追溯根因,打印根因
4. 打印日志的时候,添加trace_id

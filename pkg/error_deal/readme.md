
# err处理
如果想要实现一种方式,让error的处理变得简单一些
1. 调用方接受到error之后,如果没有额外的逻辑处理(比如要拿这个error做一些文章,做一些判断),可以将这个error抛给上层,自己不做任何处理
2. err 日志的打印不需要散落在业务代码中,统一在一个地方打印,代码中只需要[info日志]和一些[你业务上特别需要的]日志
3. error 要能区分出来是业务异常还是第三方库/基础组件 异常
4. 需要有一个专职处理error的middleware,或者在日志的middleware里面处理error
5. middleware接受到error之后,会对具体的error类型做一些判断(如果需要),打印error,要求能够打印出trace,能够追溯出根因,能够带上trace_id(链路追踪)



## 业务error
我们自己业务生成的异常,大部分情况下,业务error是由于逻辑判断产生的,比如不合法的参数,订单超时.
这些error,需要有明确的code和msg,返回给前端,给用户显示出来.
所以理论上,当产生了一个业务err,上层接受到业务error,可以直接抛给更上层,自己无需处理.
除非你的代码中需要更换文案之类的,否则理论上说,一个biz error一旦产生,直到达到middleware,中间途径中不会再产生新的error了

## 第三方库,基础组件 产生的error
一般来说,第三方库/我们自己写的基础组件,如果出现了异常,返回的一般是一个 [普通的error] ,无堆栈信息
如果这个err直接返回到middleware 他没有trace堆栈信息,所以很难溯源
比如你的middleware 收到了一个 gorm,返回record not found; 此时你无法追溯到底是哪一行给你返回了这个错误,就很吃瘪


## 项目内使用的规范
tips :不要在代码中使用err.Error()来产生一个新的error, 然后根据err字符串判断,做一些操作; 一旦这样,error文案发生变化,你的逻辑可能会失效,引发大问题

### 我们自己写的 pkg目录下的代码,产生的error

#### 简单版本
请使用
import 	"github.com/pkg/errors"
import 	"github.com/pkg/errors"
import 	"github.com/pkg/errors"
重要的事情说三遍

使用 errors.New("你的报错文案")
或者 errors.New(err)  包装一下第三方库给你返回的error

总之,pkg的所有包内,返回给 业务代码的error,必须含有堆栈信息


#### 复杂版本
先去ecode定义异常
``` go
var HealthyCheckErr = errordeal.NewToolErrWithOpt(
	errordeal.ToolErrOptionWithCode(1),
	errordeal.ToolErrOptionWithMsg("健康检查异常"),
)
```

在代码中直接返回 
```go
    return ecode.HealthyCheckErr 
```


如果有需要包装的更底层 err,可以使用
```go
    return HealthyCheckErr.AddDetail(err)
```


### 业务error使用
业务error必须预先定义,没有简单版本 hhhhh

```go
    robotOrder,err := dao.RobotOrderDao.getByID()
	
	if err != nil {
		return ecode.RobotOrderDoNotExistError.AddDetail(err)
    }
	
	if order.Status !=1 {
        return ecode.RobotOrderStatusError
    }

```




## 参考资料: kratos及grpc的要求
api层,使用krators的话, 整个api层是使用.proto文件生成的
api(pb生成) 直接调用的service func 返回给api的error,需要使用proto的&Status{} 结构
如果不这么返回的话,统一被认为是500,并且返回给前端的msg 很难看,非常不友好
详细在这段代码中  middleware logger
```go
c.Next()
err := c.Error
cerr := ecode.Cause(err)
```

```go
// Cause cause from error to ecode.
func Cause(e error) Codes {
	if e == nil {
		return OK
	}
	ec, ok := errors.Cause(e).(Codes)
	if ok {
		return ec
	}
	return String(e.Error())
}
```

pop-global-server的使用方法是
1. 定义一个业务异常,这个业务异常本质是一个 proto的&Status{}结构体
```go
	RobotOrderTooManyGoodsError        = ecode.Error(101001, "101001||The quantity of goods can not exceed 100")
```

2. 出现异常的时候,返回这个 ecode.RobotOrderTooManyGoodsError
```go
    if len(goodsInfos) > 100 {
        return robotEcode.RobotOrderTooManyGoodsError
    }
```

这样做的不足之处是,如果有更深层的异常,会被丢弃; 以及 需要单独打印日志,否则会丢失堆栈信息

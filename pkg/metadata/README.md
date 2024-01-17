# metadata
本身作用很简单
就是利用ctx来存储一些我们项目自己使用的数据
但是其中有一段
仿写自grpc的metadata

```go

// TODO(laizhongshi): 这里是我最疑问的地方
// 为什么使用一个 serverMetadataKey{} 作为key,每次NewServerContext()的时候,都去new一个struct?
// 并且FromServerContext的时候,也new一个serverMetadataKey{} 来作为key 取出metadata?
// 为什么不用常量?? 是为了节省内存吗
type serverMetadataKey struct{}

// NewServerContext creates a new context with client md attached.
func NewServerContext(ctx context.Context, md Metadata) context.Context {
	return context.WithValue(ctx, serverMetadataKey{}, md)
}

// FromServerContext returns the server metadata in ctx if it exists.
func FromServerContext(ctx context.Context) (Metadata, bool) {
	md, ok := ctx.Value(serverMetadataKey{}).(Metadata)
	return md, ok
}


```


做了一个简单测试

```go

package main

import "fmt"

type serverMetadataKey struct{}

type serverMetadataKeyNew struct{}

func main() {

	a := serverMetadataKey{}

	b := serverMetadataKey{}

	c := serverMetadataKeyNew{}

	fmt.Printf("Address of struct = %+v: %p\n", a, &a)

	fmt.Printf("Address of struct = %+v: %p\n", b, &b)

	fmt.Printf("Address of struct = %+v: %p\n", c, &c)
	
	// 三者address一致

	if a == b {
		fmt.Println("a == b")
	}
	
	// a 和 b 是完全相等的
	// a 和 c 是不能比较的
	//if a == c {
	//	
    //}

	// 测试 a和c 是否在map中一致
	d := map[interface{}]interface{}{}
	d[a] = b

	_, oka := d[a]
	_, okb := d[b]
	_, okc := d[c]

	fmt.Println(oka)
	fmt.Println(okb)
	fmt.Println(okc)

	
	// a和b 都能取出
	// c不行
	
	// 结论,所有的空结构体,不论什么type 指针地址都是一致的
	// 但是不同结构体,仍然不是一个东西

}


```

参考文
Go 最细节篇 — 空结构体是什么?
https://juejin.cn/post/6908733156707287048

本质上来讲，使用空结构体的初衷只有一个：节省内存，但是更多的情况，节省的内存其实很有限，这种情况使用空结构体的考量其实是：根本不关心结构体变量的值

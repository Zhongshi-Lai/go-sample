# ecode

本文档仅存储预先定义好的一些http_code和reason

reason 是字符串,要求全局唯一,方便追踪

定义好一些变量

预先定义的主要作用是
1. 一些复用率高的异常码,可以预先定义好,提高复用率
2. 给reason占位


这些变量仅用于存储 code,reason,msg,实际代码中,不能给grpc/http框架返回相应的ecode结构体
因为它连一个error都不是

实际使用范例请见 pkg/errors/README.md


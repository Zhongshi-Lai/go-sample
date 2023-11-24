package errordeal

import (
	"github.com/pkg/errors"
)

type ToolErr struct {
	ToolCode int
	ToolMsg  string
	InnerErr error
}

func (te *ToolErr) Error() string { return te.ToolMsg } // 符合grpc标准需要
func (te *ToolErr) Code() int {
	if te.ToolCode == 0 {
		return toolCommonErrorCode
	}
	return te.ToolCode
}
func (te *ToolErr) Message() string {
	if te.ToolMsg == "" {
		return toolCommonErrorMsg
	}
	return te.ToolMsg
}
func (te *ToolErr) Details() []interface{} { return nil }
func (te *ToolErr) DetailErr() error       { return te.InnerErr }
func (te *ToolErr) Cause() error {
	if te == nil {
		return nil
	}
	return te.DetailErr()
}
func (te *ToolErr) PkgSpecial() int { return te.Code() }
func (te *ToolErr) AddDetail(innerErr error) error {
	// 当一个toolErr需要添加内部错误的时候,使用这个方法
	// 重新初始化一个err,返回
	// 使用原先的code和msg
	return &ToolErr{
		ToolCode: te.Code(),
		ToolMsg:  te.Message(),
		InnerErr: innerErr,
	}
}

func NewSimpleToolErr(err error) error {
	// 通常情况下,直接使用这个函数即可
	// 不设定code,msg,在代码里面直接使用 NewSimpleToolErr(err)
	return &ToolErr{
		InnerErr: errors.New(err.Error()),
	}
}

func NewToolErrWithOpt(opts ...ToolErrOption) error {
	// 当你需要为toolError 设定一个code和msg的时候,使用这个函数
	// 此函数仅在ecode文件夹中使用
	// 预先定义ecode,然后业务逻辑中使用AddDetail
	e := &ToolErr{}
	for _, f := range opts {
		f(e)
	}
	// 通常情况下,直接使用这个函数即可
	return e
}

type ToolErrOption func(opt *ToolErr)

func ToolErrOptionWithCode(code int) ToolErrOption {
	return func(opt *ToolErr) {
		opt.ToolCode = code
	}
}

func ToolErrOptionWithMsg(msg string) ToolErrOption {
	return func(opt *ToolErr) {
		opt.ToolMsg = msg
	}
}

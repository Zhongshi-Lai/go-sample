package errordeal

const toolCommonErrorCode = 500001
const toolCommonErrorMsg = "服务异常,请稍后重试"

type BizError interface {
	Error() string          // 符合grpc标准需要
	Code() int              // 符合grpc标准需要
	Message() string        // 符合grpc标准需要
	Details() []interface{} // 符合grpc标准需要

	Cause() error     // 追溯根因使用
	DetailErr() error // 具体的错误,根因放置于此

	BizSpecial() int // 区分biz和pkg
}

type PkgError interface {
	Error() string          // 符合grpc标准需要
	Code() int              // 符合grpc标准需要
	Message() string        // 符合grpc标准需要
	Details() []interface{} // 符合grpc标准需要

	Cause() error     // 追溯根因使用
	DetailErr() error // 具体的错误,根因放置于此

	PkgSpecial() int // 区分biz和pkg
}

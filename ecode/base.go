package ecode

type ECode struct {
	Code    int32  // 错误码，跟 http-status 一致，并且在 grpc 中可以转换成 grpc-status
	Reason  string // 错误原因，定义为业务判定错误码
	Message string // 错误信息，为用户可读的信息，可作为用户提示内容
}

var InternalError = ECode{
	Code:    500,
	Reason:  "internal_error",
	Message: "internal server error, please try again later",
}

var DatabaseError = ECode{
	Code:    500,
	Reason:  "database_error",
	Message: "internal server error, please try again later",
}

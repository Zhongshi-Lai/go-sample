package ecode

// 仅预先定义一个reason,占位,不定义httpcode和msg,在实际撰写代码的时候,再定义http code和message
const (
	OnlyDefineReason = "only_define_reason"
)

// define all error info

var ComplatelyErrorInfo = ECode{
	Code:    500,
	Reason:  "complately_error_info",
	Message: "show user some info",
}

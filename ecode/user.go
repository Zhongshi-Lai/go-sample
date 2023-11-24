package ecode

import errordeal "go-sample/pkg/error_deal"

var UserNotExist = errordeal.BizErr{
	BizCode:    0,
	BizMessage: "",
}

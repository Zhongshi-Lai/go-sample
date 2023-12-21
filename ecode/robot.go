package ecode

import errordeal "go-sample/pkg/error_deal"

var RobotOrderDoNotExistError = errordeal.BizErr{
	BizCode:    101033,
	BizMessage: "Paypal payment created exception",
}


var RobotOrderStatusError = errordeal.BizErr{
	BizCode:    101034,
	BizMessage: "Paypal payment created exception",
}
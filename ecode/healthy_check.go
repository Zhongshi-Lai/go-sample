package ecode

import errordeal "go-sample/pkg/error_deal"

var HealthyCheckErr = errordeal.NewToolErrWithOpt(
	errordeal.ToolErrOptionWithCode(1),
	errordeal.ToolErrOptionWithMsg("健康检查异常"),
)

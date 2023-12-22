package errordeal



type BizErr struct {
	BizCode    int
	BizMessage string
	InnerErr   error
}



func (be *BizErr) Error() string          { return be.Message() }
func (be *BizErr) Code() int              { return be.BizCode }
func (be *BizErr) Message() string        { return be.BizMessage }
func (be *BizErr) Details() []interface{} { return nil }
func (be *BizErr) DetailErr() error       { return be.InnerErr }
func (be *BizErr) Cause() error {
	if be == nil {
		return nil
	}
	return be.DetailErr()
}
func (be *BizErr) BizSpecial() int { return be.Code() } // 此方法仅用于 区分 biz和pkg
func (be *BizErr) AddDetail(innerErr error) error {
	// 当一个bizError需要添加内部错误的时候,使用这个方法
	// 重新初始化一个err,返回
	// 使用原先的code和msg
	return &BizErr{
		BizCode:    be.Code(),
		BizMessage: be.Message(),
		InnerErr:   innerErr,
	}
}

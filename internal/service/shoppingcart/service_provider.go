package shoppingcart

type Service struct{}

// 根据fx in 确定调用线路图
type FxInService struct{}

func ServiceInstance(in FxInService) *Service {
	return nil
}

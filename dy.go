package dy

import "context"

var cxt = context.Background()

const (
	// AuthUrl 赋权
	AuthUrl string = "https://auth.dylk.com/auth-isv/"
	// ClientToken 获取token
	ClientToken = "https://open.douyin.com/oauth/client_token/"
	// ShopPoiQuery 查询店铺信息
	ShopPoiQuery = "https://open.douyin.com/goodlife/v1/shop/poi/query/"
	// OrderQuery 查询订单信息
	OrderQuery = "https://open.douyin.com/goodlife/v1/trade/order/query/"
	// CouponsPrepare 查询订单信息
	CouponsPrepare = "https://open.douyin.com/goodlife/v1/fulfilment/certificate/prepare/"
	// CouponsVerify 验券
	CouponsVerify = "https://open.douyin.com/goodlife/v1/fulfilment/certificate/verify/"
	// Goods 获取商品信息
	Goods = "https://open.douyin.com/goodlife/v1/goods/product/online/get/"
	// GoodsList 获取商品列表信息
	GoodsList = "https://open.douyin.com/goodlife/v1/goods/product/online/get/"
)

type DyInterface interface {
	GenAuthValidUrl(dto *GenAuthWithBindValidUrlDto) (result string, err error)
	GetCrmQuery(params map[string]string) (interface{}, error)
	GetOrder(params map[string]string) (interface{}, error)
	GetShopPoi(params map[string]string) (interface{}, error)
	CouponsVerifyPrepare(params map[string]string) (interface{}, error)
	CouponsVerifyOn(params map[string]string) (interface{}, error)
	GoodsList(params map[string]string) (interface{}, error)
	getAccessToken() (string, error)
	getHeadersToken() map[string]string
}
type Dy struct {
	Config Config
}
type Config struct {
	AppId     string
	AppSecret string
	AccountId string
}

package dy

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func (d *Dy) GetCrmQuery(params map[string]string) (interface{}, error) {
	headers := d.getHeadersToken()
	resp, err := Request(cxt, http.MethodGet, headers, params, OrderQuery)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// GetOrder 获取订单信息
//
//	params map[string]string{
//		"account_id": "",
//		"order_id":   "",
//	}
func (d *Dy) GetOrder(params map[string]string) (interface{}, error) {
	headers := d.getHeadersToken()
	resp, err := Request(cxt, http.MethodGet, headers, params, OrderQuery)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// GetShopPoi 获取店铺信息
//
//	params map[string]string{
//		"account_id": "",
//		"page":       "1",
//		"size":       "100",
//	}
func (d *Dy) GetShopPoi(params map[string]string) (interface{}, error) {
	headers := d.getHeadersToken()
	resp, err := Request(cxt, http.MethodGet, headers, params, ShopPoiQuery)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// CouponsVerifyPrepare 验券准备
//
//	params map[string]string{
//		"poi_id": "",
//		"code":   "",
//	}
func (d *Dy) CouponsVerifyPrepare(params map[string]string) (interface{}, error) {
	headers := d.getHeadersToken()
	resp, err := Request(cxt, http.MethodGet, headers, params, CouponsPrepare)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// CouponsVerifyOn 验券
//
//	params map[string]string{
//		"verify_token": "",
//		"poi_id":         "",
//	}
func (d *Dy) CouponsVerifyOn(params map[string]string) (interface{}, error) {
	headers := d.getHeadersToken()
	resp, err := Request(cxt, http.MethodGet, headers, params, CouponsVerify)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (d *Dy) GetGoods(params map[string]string) (interface{}, error) {
	headers := d.getHeadersToken()
	resp, err := Request(cxt, http.MethodGet, headers, params, Goods)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// GoodsList 获取商品列表
// params map[string]string{
// "account_id" : “”,
// "count" : 50,
// "goods_creator_type" : 1,
// "cursor"=>0
// }
func (d *Dy) GoodsList(params map[string]string) (interface{}, error) {
	headers := d.getHeadersToken()
	resp, err := Request(cxt, http.MethodGet, headers, params, GoodsList)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// GetAccessToken 获取Token
func (d *Dy) getAccessToken() (string, error) {
	params := map[string]string{
		"client_key":    d.Config.AppId,
		"client_secret": d.Config.AppSecret,
		"grant_type":    "client_credential",
	}
	resp, err := Request(cxt, http.MethodPost, map[string]string{}, params, ClientToken)
	if err != nil {
		return "", err
	}
	if respMap, ok := resp.(map[string]interface{}); ok {
		if dataMap, ok := respMap["data"].(map[string]interface{}); ok {
			return dataMap["access_token"].(string), nil
		}
	}
	return "", fmt.Errorf("resp: %v", "resp is not map[string]interface{}")

}
func (d *Dy) getHeadersToken() map[string]string {
	accessToken, _ := d.getAccessToken()
	return map[string]string{
		"content-type": "application/json",
		"access-token": accessToken,
	}
}

type GenAuthWithBindValidUrlDto struct {
	// url开始生效时间
	Timestamp string
	// 服务商自定义字符串，长度不可超过1000字节
	Extra string
	// 解决方案 具体值参照
	SolutionKey string
	// 授权能力列表
	PermissionKeys []string
	// 外部门店ID
	OutShopId string
}

func (d *Dy) GenAuthValidUrl(dto *GenAuthWithBindValidUrlDto) (result string, err error) {
	parsedUrl, err := url.Parse(AuthUrl)
	if err != nil {
		return "", err
	}

	query := map[string]string{}
	query["client_key"] = d.Config.AppId
	query["timestamp"] = dto.Timestamp
	query["charset"] = "UTF-8"
	query["solution_key"] = dto.SolutionKey
	query["permission_keys"] = strings.Join(dto.PermissionKeys, ",")
	query["out_shop_id"] = dto.OutShopId
	if dto.Extra != "" {
		query["extra"] = dto.Extra
	}

	signResult := SignV2(d.Config.AppSecret, "", query)

	// set final url params
	parsedQuery := parsedUrl.Query()
	parsedQuery.Add("sign", signResult)
	for k, v := range query {
		parsedQuery.Add(k, v)
	}

	parsedUrl.RawQuery = parsedQuery.Encode()
	return parsedUrl.String(), nil
}

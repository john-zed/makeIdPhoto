package liantuofu

import (
	"strconv"
	"time"
)

const (
	AppId        = "EW_N4249165957"
	AppKey       = "b59f43153152248c939407b60679b468"
	MerchantCode = "EW_N7174698605"
	OpenId       = "oHucvv64jHMx6hXbPRIdI2FXI_i4"
	// 快速接入版本 请求地址
	UrlQIPrecreate      = "http://api.liantuofu.com/open/precreate"    // 主扫支付
	UrlQIPayQuery       = "http://api.liantuofu.com/open/pay/query"    // 订单查询
	UrlQIPayRefund      = "http://api.liantuofu.com/open/refund"       // 订单退款
	UrlQIPayRefundQuery = "http://api.liantuofu.com/open/refund/query" // 订单退款查询
	UrlQIPayBill        = "http://api.liantuofu.com/open/bill"         // 订单账单查询

	UrlNotify = "http://staff.sxsc.ltd/liantuofuNotify" //异步通知地址

	TotalAmount float64 = 3.00 // 支付金额
)

/**
生成订单号
由商户生成的该笔交易的全局唯一ID，商户需确保其唯一性，重新发起一笔支付需要使用新订单号，避免重复支付。
后续可通过该ID查询对应订单信息。 建议值：公司简称+门店编号+时间戳+序列 支持8-64位数字、英文字母、“-”及“_”，其他字符不支持
*/
func MakeTradeNo() string {
	return "SXSC_" + MerchantCode + "_" + strconv.Itoa(int(time.Now().Local().Unix()))
}

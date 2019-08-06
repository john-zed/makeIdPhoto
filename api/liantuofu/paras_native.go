package liantuofu

import (
	"math/rand"
	"strconv"
)

func GetPrecreateReqParas(totalAmount float64, outTradeNo string) map[string]string {
	paraMap := make(map[string]string)
	paraMap["appId"] = AppId                                              // 合作方标识
	paraMap["random"] = strconv.Itoa(rand.Int())                          // 随机数
	paraMap["merchantCode"] = MerchantCode                                // 商户编号
	paraMap["outTradeNo"] = outTradeNo                                    // 订单号
	paraMap["totalAmount"] = strconv.FormatFloat(totalAmount, 'E', 2, 32) // 订单总金额
	paraMap["channel"] = "WXPAY"                                          // 支付渠道
	paraMap["tradeType"] = "NATIVE"                                       // 支付交易类型 动态二维码
	paraMap["notifyUrl"] = UrlNotify                                      // 异步通知地址
	paraMap["sign"] = CreateSign(paraMap, AppKey)                         // 签名
	return paraMap
}

func GetOrderQueryReqParas(outTradeNo string) map[string]string {
	paraMap := make(map[string]string)
	paraMap["appId"] = AppId                      // 合作方标识
	paraMap["random"] = strconv.Itoa(rand.Int())  // 随机数
	paraMap["merchantCode"] = MerchantCode        // 商户编号
	paraMap["outTradeNo"] = outTradeNo            // 订单号
	paraMap["sign"] = CreateSign(paraMap, AppKey) // 签名
	return paraMap
}

func GetOpenPayReqParas(totalAmount float64, outTradeNo string, payChannel string, openId string) map[string]string {
	paraMap := make(map[string]string)
	paraMap["appId"] = AppId                                              // 合作方标识
	paraMap["random"] = strconv.Itoa(rand.Int())                          // 随机数
	paraMap["merchantCode"] = MerchantCode                                // 商户编号
	paraMap["outTradeNo"] = outTradeNo                                    // 订单号
	paraMap["totalAmount"] = strconv.FormatFloat(totalAmount, 'E', 2, 32) // 订单总金额
	paraMap["channel"] = payChannel                                       // 支付渠道
	paraMap["tradeType"] = "JSAPI"                                        // 支付交易类型 公众号支付
	paraMap["notifyUrl"] = UrlNotify                                      // 异步通知地址
	paraMap["openId"] = openId                                            // 消费者用户标识
	paraMap["sign"] = CreateSign(paraMap, AppKey)                         // 签名
	return paraMap
}

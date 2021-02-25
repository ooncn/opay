package alipay

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/ooncn/common/obj"
	"github.com/ooncn/common/util"
	"github.com/ooncn/opay/t"
	"regexp"
)

//工具类API > 用户登陆授权
var AlipayGateway = "https://openapi.Alipay.com/gateway.do"
var HttpHeader = map[string]string{
	"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:77.0) Gecko/2010010 1 Firefox/77.0",
	"Accept":          "*/*",
	"Accept-Encoding": "gzip, deflate",
	"Connection":      "Keep-Alive",
}

func (r *Request) HTTPPOST(method string) *util.HttpOk {
	r.SetMethod(method)
	var h = util.HttpOk{Url: AlipayGateway, Method: "POST"}
	h.Params = r.ParamURL()
	h.ContentType = "application/x-www-form-urlencoded"
	h.Header = HttpHeader
	return &h
}

//user.info.auth
func (r *Request) UserInfoAuth() string {
	data := AlipayUserInfoAuth{
		Scopes: []string{"auth_base"},
		State:  util.RandAaInt(10),
	}
	r.SetMethod("user.info.auth")
	r.SetBizContent(util.JsonToStr(data))
	return r.ToUri()
} // 获取用户授权
func (r *Request) SystemOauthToken(code, grantType string) (resp ResponseAlipaySystemOauthToken, err error) {
	if grantType == "refresh_token" {
		r.Set("grant_type", "refresh_token")
		r.Set("refresh_token", code)
	} else {
		r.Set("grant_type", "authorization_code")
		r.Set("code", code)
	}
	h := r.HTTPPOST("system.oauth.token")
	h.Query()
	str := h.ResponseBody
	if str == "" {
		return
	}
	util.JsonToType(str, &resp)
	if resp.ErrorResponse != nil {
		err = errors.New(resp.ErrorResponse.SubCode + ":" + resp.ErrorResponse.SubMsg)
		return
	}

	re := resp.AlipaySystemOauthTokenResponse
	err = re.Error()
	return
} // 换取授权访问令牌
func (r *Request) UserInfoShare(authToken string) (resp AlipayUserInfoShare, err error) {
	r.Set("auth_token", authToken)
	h := r.HTTPPOST("user.info.share")
	h.Query()
	str := h.ResponseBody
	if str == "" {
		return
	}
	util.JsonToType(str, &resp)
	re := resp.AlipayUserInfoShareResponse
	err = re.Error()
	return
} // 换取授权访问令牌
func (r *Request) MiniTemplatemessageSend(data AlipayOpenAppMiniTemplatemessageSend) {
	r.SetMethod("open.app.mini.templatemessage.send")
	r.SetBizContent(util.JsonToStr(data))
	fmt.Println(r.ToUri())
} // 小程序发送模板消息
func (r *Request) MobilePublicQrcodeCreate(gotoUrl, codeType, expireSecond, showLogo, sceneId string) (resp *AlipayMobilePublicQrcodeCreate, err error) {
	data := fmt.Sprintf(`{"expireSecond":"%s","showLogo":"%s","codeType":"%s","codeInfo":{"gotoUrl":"%s","scene":{"sceneId":"%s"}}}`, expireSecond, showLogo, codeType, gotoUrl, sceneId)
	r.SetBizContent(util.JsonToStr(data))
	method := "mobile.public.qrcode.create"
	h := r.HTTPPOST(method)
	h.Query()
	str := h.ResponseBody
	if str == "" {
		err = errors.New(method + "请求失败")
		return
	}
	fmt.Println(str)
	util.JsonToType(str, &resp)
	re := resp.ErrorResponse
	err = re.Error()
	return
} // 带参推广二维码接口

func (r *Request) TradePay(outTradeNo, scene, authCode, subject string) (resp AlipayTradePay, err error) {
	var bz = make(map[string]interface{})
	bz["out_trade_no"] = outTradeNo
	bz["scene"] = scene
	bz["authCode"] = authCode
	bz["subject"] = subject
	r.SetBizContent(util.JsonToStr(bz))
	method := "trade.pay"
	h := r.HTTPPOST(method)
	h.Query()
	str := h.ResponseBody
	if str == "" {
		err = errors.New(method + "请求失败")
		return
	}
	fmt.Println(str)
	util.JsonToType(str, &resp)
	re := resp.AlipayTradePayResponse
	err = re.Error()
	return
} //统一收单交易支付接口 收银员使用扫码设备读取用户手机支付宝“付款码”/声波获取设备（如麦克风）读取用户手机支付宝的声波信息后，将二维码或条码信息/声波信息通过本接口上送至支付宝发起支付。
func (r *Request) TradePrecreate(outTradeNo, subject string, totalAmount float64) (resp AlipayTradePrecreate, err error) {
	var bz = make(map[string]interface{})
	bz["out_trade_no"] = outTradeNo
	bz["subject"] = subject
	bz["total_amount"] = fmt.Sprintf("%.2f", totalAmount)
	bz["product_code"] = "FACE_TO_FACE_PAYMENT"
	r.SetBizContent(util.JsonToStr(bz))
	method := "trade.precreate"
	h := r.HTTPPOST(method)
	h.Query()
	str := h.ResponseBody
	if str == "" {
		err = errors.New(method + "请求失败")
		return
	}
	fmt.Println(str)
	util.JsonToType(str, &resp)
	re := resp.AlipayTradePrecreateResponse
	err = re.Error()
	return
} //统一收单线下交易预创建 收银员通过收银台或商户后台调用支付宝接口，生成二维码后，展示给用户，由用户扫描二维码完成订单支付。
func (r *Request) TradeAppPay(outTradeNo, subject string, totalAmount float64) (resp AlipayTradeAppPay, err error) {
	var bz = make(map[string]interface{})
	bz["out_trade_no"] = outTradeNo
	bz["subject"] = subject
	bz["total_amount"] = fmt.Sprintf("%.2f", totalAmount)
	bz["product_code"] = "QUICK_MSECURITY_PAY"
	r.SetBizContent(util.JsonToStr(bz))
	method := "trade.app.pay"
	h := r.HTTPPOST(method)
	h.Query()
	str := h.ResponseBody
	if str == "" {
		err = errors.New(method + "请求失败")
		return
	}
	fmt.Println(str)
	util.JsonToType(str, &resp)
	re := resp.AlipayTradeAppPayResponse
	err = re.Error()
	return
} //手机网站支付接口2.0 外部商户创建订单并支付
func (r *Request) TradeWapPay(outTradeNo, subject, quitUrl string, totalAmount float64) (resp AlipayTradeWapPay, err error) {
	var bz = make(map[string]interface{})
	bz["out_trade_no"] = outTradeNo
	bz["subject"] = subject
	bz["quit_url"] = quitUrl
	bz["total_amount"] = fmt.Sprintf("%.2f", totalAmount)
	bz["product_code"] = "QUICK_WAP_WAY"
	r.SetBizContent(util.JsonToStr(bz))
	method := "trade.wap.pay"
	h := r.HTTPPOST(method)
	h.Query()
	str := h.ResponseBody
	if str == "" {
		err = errors.New(method + "请求失败")
		return
	}
	fmt.Println(str)
	util.JsonToType(str, &resp)
	re := resp.AlipayTradeWapPayResponse
	err = re.Error()
	return
} //手机网站支付接口2.0 外部商户创建订单并支付

func (r *Request) TradePagePay(outTradeNo, subject string, totalAmount float64) (resp AlipayTradePagePay, err error) {
	var bz = make(map[string]interface{})
	bz["out_trade_no"] = outTradeNo
	bz["subject"] = subject
	bz["total_amount"] = fmt.Sprintf("%.2f", totalAmount)
	bz["product_code"] = "FAST_INSTANT_TRADE_PAY"
	r.SetBizContent(util.JsonToStr(bz))
	method := "trade.page.pay"
	h := r.HTTPPOST(method)
	h.Query()
	str := h.ResponseBody
	if str == "" {
		err = errors.New(method + "请求失败")
		return
	}
	fmt.Println(str)
	util.JsonToType(str, &resp)
	re := resp.AlipayTradePagePayResponse
	err = re.Error()
	return
} //统一收单下单并支付页面接口 PC场景下单并支付

func (r *Request) TradeCreate(outTradeNo, subject string, totalAmount float64) (resp AlipayTradeCreate, err error) {
	var bz = make(map[string]interface{})
	bz["out_trade_no"] = outTradeNo
	bz["subject"] = subject
	bz["total_amount"] = fmt.Sprintf("%.2f", totalAmount)
	bz["product_code"] = "FACE_TO_FACE_PAYMENT"
	r.SetBizContent(util.JsonToStr(bz))
	method := "trade.create"
	h := r.HTTPPOST(method)
	h.Query()
	str := h.ResponseBody
	if str == "" {
		err = errors.New(method + "请求失败")
		return
	}
	fmt.Println(str)
	util.JsonToType(str, &resp)
	re := resp.AlipayTradeCreateResponse
	err = re.Error()
	return
} //统一收单交易创建接口 商户通过该接口进行交易的创建下单
func (r *Request) TradeQuery(trade_no string, t int) (resp AlipayTradeQuery, err error) {
	var bz = make(map[string]interface{})
	if t > 0 {
		bz["trade_no"] = trade_no
	} else {
		bz["out_trade_no"] = trade_no
	}
	r.SetBizContent(util.JsonToStr(bz))
	method := "trade.query"
	h := r.HTTPPOST(method)
	h.Query()
	str := h.ResponseBody
	if str == "" {
		err = errors.New(method + "请求失败")
		return
	}
	fmt.Println(str)
	util.JsonToType(str, &resp)
	re := resp.AlipayTradeQueryResponse
	err = re.Error()
	return
} //统一收单线下交易查询 该接口提供所有支付宝支付订单的查询，商户可以通过该接口主动查询订单状态，完成下一步的业务逻辑。 需要调用查询接口的情况： 当商户后台、网络、服务器等出现异常，商户系统最终未接收到支付通知； 调用支付接口后，返回系统错误或未知交易状态情况； 调用trade.pay，返回INPROCESS的状态； 调用trade.cancel之前，需确认支付状态；
func (r *Request) TradeCancel(trade_no string, t int) (resp AlipayTradeCancel, err error) {
	var bz = make(map[string]interface{})
	if t > 0 {
		bz["trade_no"] = trade_no
	} else {
		bz["out_trade_no"] = trade_no
	}
	r.SetBizContent(util.JsonToStr(bz))
	method := "trade.cancel"
	h := r.HTTPPOST(method)
	h.Query()
	str := h.ResponseBody
	if str == "" {
		err = errors.New(method + "请求失败")
		return
	}
	fmt.Println(str)
	util.JsonToType(str, &resp)
	re := resp.AlipayTradeCancelResponse
	err = re.Error()
	return
} //统一收单交易撤销接口 支付交易返回失败或支付系统超时，调用该接口撤销交易。如果此订单用户支付失败，支付宝系统会将此订单关闭；如果用户支付成功，支付宝系统会将此订单资金退还给用户。 注意：只有发生支付系统超时或者支付结果未知时可调用撤销，其他正常支付的单如需实现相同功能请调用申请退款API。提交支付交易后调用【查询订单API】，没有明确的支付结果再调用【撤销订单API】。
func (r *Request) TradeClose(trade_no, operator_id string, t int) (resp AlipayTradeClose, err error) {
	var bz = make(map[string]interface{})
	if t > 0 {
		bz["trade_no"] = trade_no
	} else {
		bz["out_trade_no"] = trade_no
	}
	bz["operator_id"] = operator_id //卖家端自定义的的操作员 ID
	r.SetBizContent(util.JsonToStr(bz))
	method := "trade.close"
	h := r.HTTPPOST(method)
	h.Query()
	str := h.ResponseBody
	if str == "" {
		err = errors.New(method + "请求失败")
		return
	}
	fmt.Println(str)
	util.JsonToType(str, &resp)
	re := resp.AlipayTradeCloseResponse
	err = re.Error()
	return
} //统一收单交易关闭接口 用于交易创建后，用户在一定时间内未进行支付，可调用该接口直接将未付款的交易进行关闭。
func (r *Request) TradeRefund(trade_no, operator_id, out_request_no, refund_reason string, refund_amount float64, t int) (resp AlipayTradeRefund, err error) {
	var bz = make(map[string]interface{})
	if t > 0 {
		bz["trade_no"] = trade_no
	} else {
		bz["out_trade_no"] = trade_no
	}
	bz["refund_amount"] = fmt.Sprintf("%.2f", refund_amount) //需要退款的金额，该金额不能大于订单金额,单位为元，支持两位小数 	200.12
	bz["operator_id"] = operator_id                          //卖家端自定义的的操作员 ID
	bz["out_request_no"] = out_request_no                    //64 	标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传。 	HZ01RF001
	bz["refund_reason"] = refund_reason                      //256 退款的原因说明

	r.SetBizContent(util.JsonToStr(bz))
	method := "trade.refund"
	h := r.HTTPPOST(method)
	h.Query()
	str := h.ResponseBody
	if str == "" {
		err = errors.New(method + "请求失败")
		return
	}
	fmt.Println(str)
	util.JsonToType(str, &resp)
	re := resp.AlipayTradeRefundResponse
	err = re.Error()
	return
} //统一收单交易退款接口 当交易发生之后一段时间内，由于买家或者卖家的原因需要退款时，卖家可以通过退款接口将支付款退还给买家，支付宝将在收到退款请求并且验证成功之后，按照退款规则将支付款按原路退到买家帐号上。 交易超过约定时间（签约时设置的可退款时间）的订单无法进行退款 支付宝退款支持单笔交易分多次退款，多次退款需要提交原支付订单的商户订单号和设置不同的退款单号。一笔退款失败后重新提交，要采用原来的退款单号。总退款金额不能超过用户实际支付金额
func (r *Request) TradePageRefund(trade_no, operator_id, out_request_no, refund_reason string, refund_amount float64,
	t int) (resp AlipayTradePageRefund, err error) {
	var bz = make(map[string]interface{})
	if t > 0 {
		bz["trade_no"] = trade_no
	} else {
		bz["out_trade_no"] = trade_no
	}
	bz["refund_amount"] = fmt.Sprintf("%.2f", refund_amount) //需要退款的金额，该金额不能大于订单金额,单位为元，支持两位小数 	200.12
	bz["operator_id"] = operator_id                          //卖家端自定义的的操作员 ID
	bz["out_request_no"] = out_request_no                    //64 	标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传。 	HZ01RF001
	bz["refund_reason"] = refund_reason                      //256 退款的原因说明

	r.SetBizContent(util.JsonToStr(bz))
	method := "trade.page.refund"
	h := r.HTTPPOST(method)
	h.Query()
	str := h.ResponseBody
	if str == "" {
		err = errors.New(method + "请求失败")
		return
	}
	fmt.Println(str)
	util.JsonToType(str, &resp)
	re := resp.AlipayTradePageRefundResponse
	err = re.Error()
	return
} //统一收单退款页面接口 当交易发生之后一段时间内，由于买家或者卖家的原因需要退款时，卖家可以通过退款页面接口将支付款退还给买家，支付宝将在收到退款请求并且验证成功之后，按照退款规则将支付款按原路退到买家帐号上。 目前该接口用于信用退款场景，通过biz_type指定信用退款。支付宝页面会提示用户退款成功或失败，退款处理完成后支付宝回跳到商户请求指定的回跳地址页面。

func (r *Request) TradeFastpayRefundQuery(trade_no, out_request_no, org_pid string,
	t int) (resp AlipayTradeFastpayRefundQuery, err error) {
	var bz = make(map[string]interface{})
	if t > 0 {
		bz["trade_no"] = trade_no
	} else {
		bz["out_trade_no"] = trade_no
	}
	bz["out_request_no"] = out_request_no //64 	请求退款接口时，传入的退款请求号，如果在退款请求时未传入，则该值为创建交易时的外部交易号
	bz["org_pid"] = org_pid               //银行间联模式下有用，其它场景请不要使用； 双联通过该参数指定需要查询的交易所属收单机构的pid;

	r.SetBizContent(util.JsonToStr(bz))
	method := "trade.fastpay.refund.query"
	h := r.HTTPPOST(method)
	h.Query()
	str := h.ResponseBody
	if str == "" {
		err = errors.New(method + "请求失败")
		return
	}
	fmt.Println(str)
	util.JsonToType(str, &resp)
	re := resp.AlipayTradeFastpayRefundQueryResponse
	err = re.Error()
	return
} //统一收单交易退款查询 商户可使用该接口查询自已通过trade.refund或trade.refund.apply提交的退款请求是否执行成功。 该接口的返回码10000，仅代表本次查询操作成功，不代表退款成功。如果该接口返回了查询数据，且refund_status为空或为REFUND_SUCCESS，则代表退款成功，如果没有查询到则代表未退款成功，可以调用退款接口进行重试。重试时请务必保证退款请求号一致。
//统一收单交易结算接口

func AlipayVerifyMap(m map[string]string, PublicKey string) error {
	sign, is := m["sign"]
	if !is {
		return errors.New("sign 不存在")
	}
	delete(m, "sign")
	str := obj.MapAndStr(m)
	return AlipayVerify(str, sign, PublicKey)
}
func AlipayVerifySyn(str, PublicKey string) (string, error) {
	reg := regexp.MustCompile(`{"alipay_.*_response":{(.*)},"sign":".*`)
	m := reg.ReplaceAllString(str, `{$1}`)
	//v := reg.FindAllStringSubmatch(str,-1)
	//fmt.Println(v[0][1])
	reg = regexp.MustCompile(`.*"sign":"([^"]*)".*`)
	sign := reg.ReplaceAllString(str, `$1`)
	//v = reg.FindAllStringSubmatch(str,-1)
	//fmt.Println(v[0][1])
	return m, AlipayVerify(m, sign, PublicKey)
}
func AlipayVerify(m, sign, PublicKey string) error {
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	err = t.RsaVerySignWithSha256([]byte(m), signBytes, []byte(fmt.Sprintf(`-----BEGIN RSA PUBLIC KEY-----
%s
-----END RSA PUBLIC KEY-----`, PublicKey)))
	return err
}

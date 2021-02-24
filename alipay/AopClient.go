package alipay

import (
	"fmt"
	"github.com/ooncn/common/obj"
	"github.com/ooncn/common/util"
	"net/url"
	"sort"
	"strings"
)

type AopClient struct {

	//应用ID
	AppId string `json:"app_id"`

	//私钥值
	RsaPrivateKey string `json:"rsa_private_key"`

	//请填写您的支付宝公钥
	RsaPublicKey string `json:"rsa_public_key"`
	//请填写您的支付宝公钥
	PublicKey string `json:"public_key"`

	//网关
	GatewayUrl string `json:"gateway_url"` //= "https://openapi.Alipay.com/gateway.do";
	//返回数据格式
	Format string `json:"format"` // "json";
	//api版本
	ApiVersion string `json:"api_version"` // "1.0";

	// 表单提交字符集编码
	PostCharset string `json:"post_charset"` // "UTF-8";

	DebugInfo bool `json:"debug_info"` // false;

	FileCharset string `json:"file_charset"` // "UTF-8";

	RESPONSE_SUFFIX string `json:"response_suffix"` // "_response";

	ERROR_RESPONSE string `json:"error_response"` // "error_response";

	SIGN_NODE_NAME string `json:"sign_node_name"` // "sign";

	//加密XML节点名称
	ENCRYPT_XML_NODE_NAME string `json:"encrypt_xml_node_name"` // "response_encrypted";

	NeedEncrypt bool `json:"need_encrypt"` // false;

	//签名类型
	SignType string `json:"sign_type"` // "RSA";

	//加密密钥和类型

	EncryptKey  string `json:"encrypt_key"`
	EncryptType string `json:"encrypt_type"` //"AES";

	TargetServiceUrl string `json:"target_service_url"` // "";

	Rsa2 *util.Rsa2
}

func (a *AopClient) NewRsa2() (err error) {
	r, err := util.NewRsa2PKCS1(a.RsaPublicKey, a.RsaPrivateKey)
	if err != nil {
		fmt.Println("支付宝加密组件失败" + err.Error())
		return err
	} else {
		a.Rsa2 = r
	}
	return
}

/**

 */
func (r *Request) LoginUri(uri, scope, state string) string {
	return fmt.Sprintf("https://openauth.alipay.com/oauth2/publicAppAuthorize.htm?app_id=%s&scope=%s&redirect_uri=%s&state=%s", r.Get("app_id"), scope, url.QueryEscape(uri), state)
} // 网页登录连接

type Request struct {
	obj.Map
	Client *AopClient
}

func NewRequest() *Request {
	if Client == nil {
		fmt.Println("支付宝配置文件缺失")
		return nil
	}
	request := make(obj.Map)
	request["app_id"] = Client.AppId
	request["format"] = Client.Format
	//request["method"] = "" 接口名称 	alipay.user.info.auth
	//request["return_url"] = "" //HTTP/HTTPS开头字符串 	https://m.alipay.com/Gk8NF23
	request["charset"] = Client.PostCharset //请求使用的编码格式，如utf-8,gbk,gb2312等 	utf-8
	request["sign_type"] = Client.SignType  //商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA2 	RSA2
	request["version"] = Client.ApiVersion  //调用的接口版本，固定为：1.0 	1.0
	//request["app_auth_token"] = "" //详见应用授权概述
	request["timestamp"] = util.TimeUtil.Sep("yyyy-MM-dd HH:mm:ss") //发送请求的时间，格式"yyyy-MM-dd HH:mm:ss" 	2014-07-24 03:07:50
	//request["biz_content"] = "util.TimeUtil.Sep("yyyy-MM-dd HH:mm:ss")" //请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
	var r Request
	r.Map = request
	return &r
}
func NewRequestClient(client AopClient) *Request {
	err := client.NewRsa2()
	if err != nil {
		fmt.Println("加密组件错误" + err.Error())
		return nil
	}
	request := make(obj.Map)
	request["app_id"] = client.AppId
	request["format"] = client.Format
	//request["method"] = "" 接口名称 	alipay.user.info.auth
	//request["return_url"] = "" //HTTP/HTTPS开头字符串 	https://m.alipay.com/Gk8NF23
	request["charset"] = client.PostCharset //请求使用的编码格式，如utf-8,gbk,gb2312等 	utf-8
	request["sign_type"] = client.SignType  //商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA2 	RSA2
	request["version"] = client.ApiVersion  //调用的接口版本，固定为：1.0 	1.0
	//request["app_auth_token"] = "" //详见应用授权概述
	request["timestamp"] = util.TimeUtil.Sep("yyyy-MM-dd HH:mm:ss") //发送请求的时间，格式"yyyy-MM-dd HH:mm:ss" 	2014-07-24 03:07:50
	//request["biz_content"] = "util.TimeUtil.Sep("yyyy-MM-dd HH:mm:ss")" //请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
	var r Request
	r.Map = request
	r.Client = &client
	return &r
}
func (r *Request) SetMethod(method string) {
	r.Set("method", method)
}
func (r *Request) SetReturnUrl(return_url string) {
	r.Set("return_url", return_url)
}
func (r *Request) SetNotifyUrl(notify_url string) {
	r.Set("notify_url", notify_url)
}
func (r *Request) SetTimestamp(timestamp string) {
	r.Set("timestamp", timestamp)
}
func (r *Request) SetBizContent(biz_content string) {
	r.Set("biz_content", biz_content)
}
func (r *Request) SetBizContentJson(biz_content interface{}) {
	str := util.JsonToStr(biz_content)
	str, _ = util.JsonClear(str)
	r.Set("biz_content", str)
}
func (r *Request) SetSign() {

}
func (r *Request) ToUri() string {
	m := r.Map
	m["sign"] = url.QueryEscape(r.SignGet())
	var data []string
	for k, v := range m {
		data = append(data, fmt.Sprintf(`%s=%s`, k,
			//url.QueryEscape(v),
			v))
	}
	sort.Strings(data)
	p := strings.Join(data, "&")
	return GATEWAY + "?" + p
}
func (r *Request) ParamURL() string {
	m := r.Map
	m["sign"] = url.QueryEscape(r.SignGet())
	var data []string
	for k, v := range m {
		data = append(data, fmt.Sprintf(`%s=%s`, k,
			//url.QueryEscape(v)))
			v))
	}
	return strings.Join(data, "&")
}
func (r *Request) SignGet() (sign string) {
	m := r.Map
	delete(m, "sign")
	str := obj.MapAndStr(m)
	sign, _ = r.Client.Rsa2.Sign(str)
	return
}
func (r *Request) CheckSign() bool {
	m := r.Map
	sign, is := m["sign"]
	if !is {
		fmt.Println("sign 不存在")
		return false
	}
	delete(m, "sign_type")
	delete(m, "sign")
	str := obj.MapAndStr(m)
	fmt.Println(str)
	err := r.Client.Rsa2.Verify(str, sign)
	fmt.Println(err)
	return err == nil
}

/*
其他 APP 或 外部 H5 跳转小程序目前有两种方式可以跳转：

Scheme 拼接方式：加上 https://ds.alipay.com/?scheme=前缀  ，
后面拼接 alipays://platformapi/startapp?appId=xxx&page=x/y/z&query=xx%3dxx
*/

type AccessParams struct {
	channel string //必填 	20 	目前支持以下值： 1. ALIPAYAPP （钱包h5页面签约） 2. QRCODE(扫码签约) 3. QRCODEORSMS(扫码签约或者短信签约) 	ALIPAYAPP
} //请按当前接入的方式进行填充，且输入值必须为文档中的参数取值范围。

type RoyaltyInfo struct {
	royalty_type         string                //150 	分账类型 卖家的分账类型，目前只支持传入ROYALTY（普通分账类型）。 	ROYALTY
	royalty_detail_infos *[]RoyaltyDetailInfos //必填 	2500 	分账明细的信息，可以描述多条分账指令，json数组。
} //描述分账信息，json格式，详见分账参数说明
type RoyaltyDetailInfos struct {
	serial_no         int     //9 	分账序列号，表示分账执行的顺序，必须为正整数 	1
	trans_in_type     string  //24 	接受分账金额的账户类型：  userId：支付宝账号对应的支付宝唯一用户号。  bankIndex：分账到银行账户的银行编号。目前暂时只支持分账到一个银行编号。 storeId：分账到门店对应的银行卡编号。 默认值为userId。 	userId
	batch_no          string  //必填 	32 	分账批次号 分账批次号。 目前需要和转入账号类型为bankIndex配合使用。 	123
	out_relation_id   string  //64 	商户分账的外部关联号，用于关联到每一笔分账信息，商户需保证其唯一性。 如果为空，该值则默认为“商户网站唯一订单号+分账序列号” 	20131124001
	trans_out_type    string  //必填 	24 	要分账的账户类型。 目前只支持userId：支付宝账号对应的支付宝唯一用户号。默认值为userId。 	userId
	trans_out         string  //必填 	16 	如果转出账号类型为userId，本参数为要分账的支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字。 	2088101126765726
	trans_in          string  //必填 	28 	如果转入账号类型为userId，本参数为接受分账金额的支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字。  如果转入账号类型为bankIndex，本参数为28位的银行编号（商户和支付宝签约时确定）。 如果转入账号类型为storeId，本参数为商户的门店ID。 	2088101126708402
	amount            float64 //必填 	9 	分账的金额，单位为元 	0.1
	desc              string  //1000 	分账描述信息 	分账测试1
	amount_percentage string  //3 	分账的比例，值为20代表按20%的比例分账 	100
}

type SignMerchantParams struct {
	sub_merchant_id                  string            //20 	子商户的商户id 	2088123412341234
	sub_merchant_name                string            //50 	子商户的商户名称 	滴滴出行
	sub_merchant_service_name        string            //50 	子商户的服务名称 	滴滴出行免密支付
	sub_merchant_service_description string            //150 	子商户的服务描述 	免密付车费，单次最高500
	period_rule_params               *PeriodRuleParams //周期管控规则参数period_rule_params，在签约周期扣款产品（如CYCLE_PAY_AUTH_P）时必传，在签约其他产品时无需传入。 周期扣款产品，会按照这里传入的参数提示用户，并对发起扣款的时间、金额、次数等做相应限制。
} //此参数用于传递子商户信息，无特殊需求时不用关注。目前商户代扣、海外代扣、淘旅行信用住产品支持传入该参数（在销售方案中“是否允许自定义子商户信息”需要选是）。
type SignParams struct {
	personal_product_code string              //必填 	64 	个人签约产品码，商户和支付宝签约时确定。 	CYCLE_PAY_AUTH_P
	sign_scene            string              //必填 	64 	协议签约场景，商户和支付宝签约时确定，商户可咨询技术支持。 	INDUSTRY|DIGITAL_MEDIA
	external_agreement_no string              //32 	商户签约号，代扣协议中标示用户的唯一签约号（确保在商户系统中唯一）。 格式规则：支持大写小写字母和数字，最长32位。 商户系统按需传入，如果同一用户在同一产品码、同一签约场景下，签订了多份代扣协议，那么需要指定并传入该值。 	test20190701
	external_logon_id     string              //100 	用户在商户网站的登录账号，用于在签约页面展示，如果为空，则不展示 	13852852877
	access_params         *AccessParams       //必填 		请按当前接入的方式进行填充，且输入值必须为文档中的参数取值范围。
	sub_merchant          *SignMerchantParams //此参数用于传递子商户信息，无特殊需求时不用关注。目前商户代扣、海外代扣、淘旅行信用住产品支持传入该参数（在销售方案中“是否允许自定义子商户信息”需要选是）。
	allow_huazhi_degrade  string              //10 	是否允许花芝GO降级成原代扣（即销售方案指定的代扣产品），在花芝GO场景下才会使用该值。取值：true-允许降级，false-不允许降级。默认为true。 	false
	sign_notify_url       string              //3000 	签约成功后商户用于接收异步通知的地址。如果不传入，签约与支付的异步通知都会发到外层notify_url参数传入的地址；如果外层也未传入，签约与支付的异步通知都会发到商户appid配置的网关地址。 	http://www.merchant.com/receiveSignNotify
} //签约参数。如果希望在sdk中支付并签约，需要在这里传入签约信息。

type PeriodRuleParams struct {
	period_type    string  //必填 	20 	周期类型period_type是周期扣款产品必填，枚举值为DAY和MONTH。 DAY即扣款周期按天计，MONTH代表扣款周期按自然月。 与另一参数period组合使用确定扣款周期，例如period_type为DAY，period=30，则扣款周期为30天；period_type为MONTH，period=3，则扣款周期为3个自然月。 自然月是指，不论这个月有多少天，周期都计算到月份中的同一日期。例如1月3日到2月3日为一个自然月，1月3日到4月3日为三个自然月。注意周期类型使用MONTH的时候，计划扣款时间execute_time不允许传28日之后的日期（可以传28日），以此避免有些月份可能不存在对应日期的情况。 	DAY
	period         int     //必填 	8 	周期数period是周期扣款产品必填。与另一参数period_type组合使用确定扣款周期，例如period_type为DAY，period=90，则扣款周期为90天。 	3
	execute_time   string  //必填 	16 	首次执行时间execute_time是周期扣款产品必填，即商户发起首次扣款的时间。精确到日，格式为yyyy-MM-dd 结合其他必填的扣款周期参数，会确定商户以后的扣款计划。发起扣款的时间需符合这里的扣款计划。 	2019-01-23
	single_amount  float64 //必填 	32 	单次扣款最大金额single_amount是周期扣款产品必填，即每次发起扣款时限制的最大金额，单位为元。商户每次发起扣款都不允许大于此金额。 	10.99
	total_amount   float64 //32 	总金额限制，单位为元。如果传入此参数，商户多次扣款的累计金额不允许超过此金额。 	600
	total_payments int64   //可选 	8 	总扣款次数。如果传入此参数，则商户成功扣款的次数不能超过此次数限制（扣款失败不计入）。 	12
} //周期管控规则参数period_rule_params，在签约周期扣款产品（如CYCLE_PAY_AUTH_P）时必传，在签约其他产品时无需传入。 周期扣款产品，会按照这里传入的参数提示用户，并对发起扣款的时间、金额、次数等做相应限制。

type SubMerchant struct {
	merchant_id   string //必填 	16 	间连受理商户的支付宝商户编号，通过间连商户入驻后得到。间连业务下必传，并且需要按规范传递受理商户编号。 	2088000603999128
	merchant_type string //32 	商户id类型， 	alipay: 支付宝分配的间连商户编号, merchant: 商户端的间连商户编号
}
type SettleInfo struct {
	settle_detail_infos *[]SettleDetailInfo //必填 	10 	结算详细信息，json数组，目前只支持一条。
	settle_period_time  string              //10 	该笔订单的超期自动确认结算时间，到达期限后，将自动确认结算。此字段只在签约账期结算模式时有效。取值范围：1d～365d。d-天。 该参数数值不接受小数点。 	7d
} //描述结算信息，json格式，详见结算参数说明
type SettleDetailInfo struct {
	trans_in_type      string  // 	必填 	64 	结算收款方的账户类型。 cardAliasNo：结算收款方的银行卡编号; userId：表示是支付宝账号对应的支付宝唯一用户号; loginName：表示是支付宝登录号； defaultSettle：表示结算到商户进件时设置的默认结算账号，结算主体为门店时不支持传defaultSettle； 	cardAliasNo
	trans_in           string  //必填 	64 	结算收款方。当结算收款方类型是cardAliasNo时，本参数为用户在支付宝绑定的卡编号；结算收款方类型是userId时，本参数为用户的支付宝账号对应的支付宝唯一用户号，以2088开头的纯16位数字；当结算收款方类型是loginName时，本参数为用户的支付宝登录号；当结算收款方类型是defaultSettle时，本参数不能传值，保持为空。 	A0001
	summary_dimension  string  //64 	结算汇总维度，按照这个维度汇总成批次结算，由商户指定。 目前需要和结算收款方账户类型为cardAliasNo配合使用 	A0001
	settle_entity_id   string  //64 	结算主体标识。当结算主体类型为SecondMerchant时，为二级商户的SecondMerchantID；当结算主体类型为Store时，为门店的外标。 	2088xxxxx;ST_0001
	settle_entity_type string  //32 	结算主体类型。 二级商户:SecondMerchant;商户或者直连商户门店:Store 	SecondMerchant、Store
	amount             float64 //必填 	9 	结算的金额，单位为元。目前必须和交易金额相同 	0.1
} //结算详细信息
type InvoiceInfo struct {
	key_info *InvoiceKeyInfo //开票关键信息
	details  string          //必填 	400 	开票内容 注：json数组格式 	[{"code":"100294400","name":"服饰","num":"2","sumPrice":"200.00","taxRate":"6%"}]
} //开票信息
type InvoiceKeyInfo struct {
	is_support_invoice    bool   //必填 	5 	该交易是否支持开票 	true
	invoice_merchant_name string //必填 	80 	开票商户名称：商户品牌简称|商户门店简称 	ABC|003
	tax_num               string //必填 	30 	税号 	1464888883494
} //开票关键信息

type TradePay struct {
	OutTradeNo string `json:"out_trade_no"` //必选 	64 	商户订单号,64个字符以内、可包含字母、数字、下划线；需保证在商户端不重复 	20150320010101001
	Scene      string `json:"scene"`        //必选 	32 	支付场景 条码支付，取值：bar_code 声波支付，取值：wave_code 	bar_code
	AuthCode   string `json:"auth_code"`    //必选 	64 	支付授权码，25~30开头的长度为16~24位的数字，实际字符串长度以开发者获取的付款码长度为准 	28763443825664394
	Subject    string `json:"subject"`      //必选 	256 订单标题 	Iphone6 16G

	ProductCode         string  //64 	销售产品码 	FACE_TO_FACE_PAYMENT
	BuyerId             string  // 	28 	买家的支付宝用户 id，如果为空，会从传入的码值信息中获取买家 ID 	2088202954065786
	seller_id           string  //28 	如果该值为空，则默认为商户签约账号对应的支付宝用户ID 	2088102146225135
	total_amount        float64 // 	11 	订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 如果同时传入【可打折金额】和【不可打折金额】，该参数可以不用传入； 如果同时传入了【可打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【订单总金额】=【可打折金额】+【不可打折金额】 	88.88
	trans_currency      string  //8 	标价币种, total_amount 对应的币种单位。支持英镑：GBP、港币：HKD、美元：USD、新加坡元：SGD、日元：JPY、加拿大元：CAD、澳元：AUD、欧元：EUR、新西兰元：NZD、韩元：KRW、泰铢：THB、瑞士法郎：CHF、瑞典克朗：SEK、丹麦克朗：DKK、挪威克朗：NOK、马来西亚林吉特：MYR、印尼卢比：IDR、菲律宾比索：PHP、毛里求斯卢比：MUR、以色列新谢克尔：ILS、斯里兰卡卢比：LKR、俄罗斯卢布：RUB、阿联酋迪拉姆：AED、捷克克朗：CZK、南非兰特：ZAR、人民币：CNY 	USD
	settle_currency     string  //8 	商户指定的结算币种，支持英镑：GBP、港币：HKD、美元：USD、新加坡元：SGD、日元：JPY、加拿大元：CAD、澳元：AUD、欧元：EUR、新西兰元：NZD、韩元：KRW、泰铢：THB、瑞士法郎：CHF、瑞典克朗：SEK、丹麦克朗：DKK、挪威克朗：NOK、马来西亚林吉特：MYR、印尼卢比：IDR、菲律宾比索：PHP、毛里求斯卢比：MUR、以色列新谢克尔：ILS、斯里兰卡卢比：LKR、俄罗斯卢布：RUB、阿联酋迪拉姆：AED、捷克克朗：CZK、南非兰特：ZAR、人民币：CNY 	USD
	discountable_amount float64 //11 	参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]。 如果该值未传入，但传入了【订单总金额】和【不可打折金额】，则该值默认为【订单总金额】-【不可打折金额】 	8.88

	body string //String 	可选 	128 	订单描述 	Iphone6 16G

} //统一收单交易支付接口 收银员使用扫码设备读取用户手机支付宝“付款码”/声波获取设备（如麦克风）读取用户手机支付宝的声波信息后，将二维码或条码信息/声波信息通过本接口上送至支付宝发起支付。
type TradeAppPay struct {
	Subject     string  `json:"subject"`      //必选 256 商品的标题/交易标题/订单标题/订单关键字等。
	OutTradeNo  string  `json:"out_trade_no"` //必选 64 	商户网站唯一订单号
	TotalAmount float64 `json:"total_amount"` //必选 9 	订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	ProductCode string  `json:"product_code"` //必选 64 	销售产品码，商家和支付宝签约的产品码 QUICK_MSECURITY_PAY

	TimeoutExpress     string        `json:"timeout_express"`      //6 	该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0 点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。
	Body               string        `json:"body"`                 //128 对一笔交易的具体描述信息。如果是多种商品，请将商品描述字符串累加传给body。
	TimeExpire         string        `json:"time_expire"`          //32 	绝对超时时间，格式为yyyy-MM-dd HH:mm。
	GoodsType          string        `json:"goods_type"`           //2 	商品主类型 :0-虚拟类商品,1-实物类商品
	PromoParams        string        `json:"promo_params"`         //512 	优惠参数 注：仅与支付宝协商后可用 	{"storeIdType":"1"}
	PassbackParams     string        `json:"passback_params"`      //512 	公用回传参数，如果请求时传递了该参数，则返回给商户时会回传该参数。支付宝只会在同步返回（包括跳转回商户网站）和异步通知时将该参数原样返回。本参数必须进行UrlEncode之后才可以发送给支付宝。
	ExtendParams       *ExtendParams `json:"extend_params"`        //业务扩展参数
	MerchantOrderNo    string        `json:"merchant_order_no"`    //32 	商户原始订单号，最大长度限制32位 	20161008001
	EnablePayChannels  string        `json:"enable_pay_channels"`  //128 	可用渠道，用户只能在指定渠道范围内支付 当有多个渠道时用“,”分隔 注，与disable_pay_channels互斥 	pcredit, moneyFund,debitCardExpress
	StoreId            string        `json:"store_id"`             //32 	商户门店编号 	NJ_001
	SpecifiedChannel   string        `json:"specified_channel"`    //128 	指定渠道，目前仅支持传入pcredit 若由于用户原因渠道不可用，用户可选择是否用其他渠道支付。 注：该参数不可与花呗分期参数同时传入
	DisablePayChannels string        `json:"disable_pay_channels"` //128 	禁用渠道，用户不可用指定渠道支付 当有多个渠道时用“,”分隔 注，与enable_pay_channels互斥 	pcredit, moneyFund,debitCardExpress
	ExtUserInfo        *ExtUserInfo  `json:"ext_user_info"`        //外部指定买家
	GoodsDetail        *GoodsDetail  `json:"goods_detail"`         //订单包含的商品列表信息，json格式，其它说明详见商品明细说明
	BusinessParams     string        `json:"business_params"`      //512 	商户传入业务信息，具体值要和支付宝约定，应用于安全，营销等参数直传场景，格式为json格式 	{"data":"123"}

	agreement_sign_params *SignParams //签约参数。如果希望在sdk中支付并签约，需要在这里传入签约信息。
} //app支付接口2.0 外部商户APP唤起快捷SDK创建订单并支付
type TradeWapPay struct {
	Subject     string  `json:"subject"`      //必选 256	商品的标题/交易标题/订单标题/订单关键字等。
	OutTradeNo  string  `json:"out_trade_no"` //必选 64 	商户网站唯一订单号
	TotalAmount float64 `json:"total_amount"` //必选 9 	订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	QuitUrl     string  `json:"quit_url"`     //必选 400 用户付款中途退出返回商户网站的地址
	ProductCode string  `json:"product_code"` //必选 64 	销售产品码，商家和支付宝签约的产品码 QUICK_WAP_WAY

	Body               string        `json:"body"`                 //128 对一笔交易的具体描述信息。如果是多种商品，请将商品描述字符串累加传给body。
	TimeoutExpress     string        `json:"timeout_express"`      //6 	该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0 点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。
	TimeExpire         string        `json:"time_expire"`          //32 	绝对超时时间，格式为yyyy-MM-dd HH:mm。
	AuthToken          string        `json:"auth_token"`           //40 	针对用户授权接口，获取用户相关数据时，用于标识用户授权关系
	GoodsType          string        `json:"goods_type"`           //2 	商品主类型 :0-虚拟类商品,1-实物类商品
	PassbackParams     string        `json:"passback_params"`      //512 	公用回传参数，如果请求时传递了该参数，则返回给商户时会回传该参数。支付宝只会在同步返回（包括跳转回商户网站）和异步通知时将该参数原样返回。本参数必须进行UrlEncode之后才可以发送给支付宝。
	PromoParams        string        `json:"promo_params"`         //512 	优惠参数 注：仅与支付宝协商后可用 	{"storeIdType":"1"}
	ExtendParams       *ExtendParams `json:"extend_params"`        //业务扩展参数
	MerchantOrderNo    string        `json:"merchant_order_no"`    //32 	商户原始订单号，最大长度限制32位 	20161008001
	EnablePayChannels  string        `json:"enable_pay_channels"`  //128 	可用渠道，用户只能在指定渠道范围内支付 当有多个渠道时用“,”分隔 注，与disable_pay_channels互斥 	pcredit, moneyFund,debitCardExpress
	DisablePayChannels string        `json:"disable_pay_channels"` //128 	禁用渠道，用户不可用指定渠道支付 当有多个渠道时用“,”分隔 注，与enable_pay_channels互斥 	pcredit, moneyFund,debitCardExpress
	StoreId            string        `json:"store_id"`             //32 	商户门店编号 	NJ_001
	GoodsDetail        *GoodsDetail  `json:"goods_detail"`         //订单包含的商品列表信息，json格式，其它说明详见商品明细说明
	SpecifiedChannel   string        `json:"specified_channel"`    //128 	指定渠道，目前仅支持传入pcredit 若由于用户原因渠道不可用，用户可选择是否用其他渠道支付。 注：该参数不可与花呗分期参数同时传入
	BusinessParams     string        `json:"business_params"`      //512 	商户传入业务信息，具体值要和支付宝约定，应用于安全，营销等参数直传场景，格式为json格式 	{"data":"123"}
	ExtUserInfo        *ExtUserInfo  `json:"ext_user_info"`        //外部指定买家
} //手机网站支付接口2.0 外部商户创建订单并支付
type TradePagePay struct {
	OutTradeNo  string  `json:"out_trade_no"` //必选 64 	商户网站唯一订单号
	ProductCode string  `json:"product_code"` //必选 64 	销售产品码，与支付宝签约的产品码名称。 注：目前仅支持FAST_INSTANT_TRADE_PAY 	FAST_INSTANT_TRADE_PAY
	TotalAmount float64 `json:"total_amount"` //必选 9 	订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	Subject     string  `json:"subject"`      //必选 256	商品的标题/交易标题/订单标题/订单关键字等。

	Body               string        `json:"body"`            //128 对一笔交易的具体描述信息。如果是多种商品，请将商品描述字符串累加传给body。
	TimeExpire         string        `json:"time_expire"`     //32 	绝对超时时间，格式为yyyy-MM-dd HH:mm。
	GoodsDetail        *GoodsDetail  `json:"goods_detail"`    //订单包含的商品列表信息，json格式，其它说明详见商品明细说明
	PassbackParams     string        `json:"passback_params"` //512 	公用回传参数，如果请求时传递了该参数，则返回给商户时会回传该参数。支付宝只会在同步返回（包括跳转回商户网站）和异步通知时将该参数原样返回。本参数必须进行UrlEncode之后才可以发送给支付宝。
	ExtendParams       *ExtendParams `json:"extend_params"`   //业务扩展参数
	GoodsType          string        `json:"goods_type"`      //2 	商品主类型 :0-虚拟类商品,1-实物类商品
	TimeoutExpress     string        `json:"timeout_express"` //6 	该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0 点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。
	PromoParams        string        `json:"promo_params"`    //512 	优惠参数 注：仅与支付宝协商后可用 	{"storeIdType":"1"}
	RoyaltyInfo        *RoyaltyInfo  `json:"royalty_info"`    //描述分账信息，json格式，详见分账参数说明
	SubMerchant        *SubMerchant
	MerchantOrderNo    string       `json:"merchant_order_no"`    //32 	商户原始订单号，最大长度限制32位 	20161008001
	EnablePayChannels  string       `json:"enable_pay_channels"`  //128 	可用渠道，用户只能在指定渠道范围内支付 当有多个渠道时用“,”分隔 注，与disable_pay_channels互斥 	pcredit, moneyFund,debitCardExpress
	StoreId            string       `json:"store_id"`             //32 	商户门店编号 	NJ_001
	DisablePayChannels string       `json:"disable_pay_channels"` //128 	禁用渠道，用户不可用指定渠道支付 当有多个渠道时用“,”分隔 注，与enable_pay_channels互斥 	pcredit, moneyFund,debitCardExpress
	qr_pay_mode        string       //2 	PC扫码支付的方式，支持前置模式和 跳转模式。前置模式是将二维码前置到商户的订单确认页的模式。需要商户在自己的页面中以 iframe 方式请求支付宝页面。具体分为以下几种：0：订单码-简约前置模式，对应 iframe 宽度不能小于600px，高度不能小于300px；1：订单码-前置模式，对应iframe 宽度不能小于 300px，高度不能小于600px；3：订单码-迷你前置模式，对应 iframe 宽度不能小于 75px，高度不能小于75px；4：订单码-可定义宽度的嵌入式二维码，商户可根据需要设定二维码的大小。跳转模式下，用户的扫码界面是由支付宝生成的，不在商户的域名下。2：订单码-跳转模式
	qrcode_width       int          //4 	商户自定义二维码宽度 注：qr_pay_mode=4时该参数生效 	100
	settle_info        *SettleInfo  //描述结算信息，json格式，详见结算参数说明
	invoice_info       *InvoiceInfo //开票信息

	SpecifiedChannel string `json:"specified_channel"` //128 	指定渠道，目前仅支持传入pcredit 若由于用户原因渠道不可用，用户可选择是否用其他渠道支付。 注：该参数不可与花呗分期参数同时传入

	AgreementSignParams *AgreementSignParams `json:"agreement_sign_params"` //签约参数。如果希望在sdk中支付并签约，需要在这里传入签约信息。
	IntegrationType     string               `json:"integration_type"`      //16 	请求后页面的集成方式。取值范围： 1. ALIAPP：支付宝钱包内 2. PCWEB：PC端访问 默认值为PCWEB。
	RequestFromUrl      string               `json:"request_from_url"`      //256 	请求来源地址。如果使用ALIAPP的集成方式，用户中途取消支付会返回该地址。 	https://
	BusinessParams      string               `json:"business_params"`       //512 	商户传入业务信息，具体值要和支付宝约定，应用于安全，营销等参数直传场景，格式为json格式 	{"data":"123"}
	ExtUserInfo         *ExtUserInfo         `json:"ext_user_info"`         //外部指定买家
} //统一收单下单并支付页面接口 PC场景下单并支付

type TradeCreate struct {
	OutTradeNo         string  `json:"out_trade_no"`        //必选 	64 	商户订单号,64个字符以内、只能包含字母、数字、下划线；需保证在商户端不重复 	20150320010101001
	TotalAmount        float64 `json:"total_amount"`        //必选 9 	订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 如果同时传入了【打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【订单总金额】=【打折金额】+【不可打折金额】 	88.88
	Subject            string  `json:"subject"`             //必选 	256 	订单标题 	Iphone6 16G
	ProductCode        string  `json:"product_code"`        //可选 	64 	销售产品码。 如果签约的是当面付快捷版，则传OFFLINE_PAYMENT 其它支付宝当面付产品传FACE_TO_FACE_PAYMENT； 不传默认使用FACE_TO_FACE_PAYMENT；    FACE_TO_FACE_PAYMENT
	SellerId           string  `json:"seller_id"`           //28 	卖家支付宝用户ID。 如果该值为空，则默认为商户签约账号对应的支付宝用户ID 	2088102146225135
	DiscountableAmount float64 `json:"discountable_amount"` //9 	可打折金额. 参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 如果该值未传入，但传入了【订单总金额】，【不可打折金额】则该值默认为【订单总金额】-【不可打折金额】 	8.88
} //统一收单交易创建接口 商户通过该接口进行交易的创建下单

type TradeQuery struct {
	OutTradeNo string `json:"out_trade_no"` //特殊可选 	64 	订单支付时传入的商户订单号,和支付宝交易号不能同时为空。 trade_no,out_trade_no如果同时存在优先取trade_no 	20150320010101001
	TradeNo    string `json:"trade_no"`     //特殊可选	64 	支付宝交易号，和商户订单号不能同时为空 	2014112611001004680 073956707
	//OrgPid       string    `json:"org_pid"`       //16 	银行间联模式下有用，其它场景请不要使用； 双联通过该参数指定需要查询的交易所属收单机构的pid; 	2088101117952222
	QueryOptions *[]string `json:"query_options"` //10 	查询选项，商户通过上送该字段来定制查询返回信息 	TRADE_SETTLE_INFO
} //统一收单线下交易查询 该接口提供所有支付宝支付订单的查询，商户可以通过该接口主动查询订单状态，完成下一步的业务逻辑。 需要调用查询接口的情况： 当商户后台、网络、服务器等出现异常，商户系统最终未接收到支付通知； 调用支付接口后，返回系统错误或未知交易状态情况； 调用alipay.trade.pay，返回INPROCESS的状态； 调用alipay.trade.cancel之前，需确认支付状态；
type TradeCancel struct {
	OutTradeNo string `json:"out_trade_no"` //特殊可选 	64 	订单支付时传入的商户订单号,和支付宝交易号不能同时为空。 trade_no,out_trade_no如果同时存在优先取trade_no 	20150320010101001
	TradeNo    string `json:"trade_no"`     //特殊可选	64 	支付宝交易号，和商户订单号不能同时为空 	2014112611001004680 073956707
} //统一收单线下交易查询 该接口提供所有支付宝支付订单的查询，商户可以通过该接口主动查询订单状态，完成下一步的业务逻辑。 需要调用查询接口的情况： 当商户后台、网络、服务器等出现异常，商户系统最终未接收到支付通知； 调用支付接口后，返回系统错误或未知交易状态情况； 调用alipay.trade.pay，返回INPROCESS的状态； 调用alipay.trade.cancel之前，需确认支付状态；

/*
{
"app_id":"",//是 	32 	支付宝分配给开发者的应用ID 	2014072300007148
"method":"",//是 	128 	接口名称 	alipay.trade.wap.pay
"format":"",//否 	40 	仅支持JSON 	JSON
"return_url":"",//否 	256 	HTTP/HTTPS开头字符串 	https://m.alipay.com/Gk8NF23
"charset":"",//是 	10 	请求使用的编码格式，如utf-8,gbk,gb2312等 	utf-8
"sign_type":"",//是 	10 	商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA2 	RSA2
"sign":"",//是 	344 	商户请求参数的签名串，详见签名 	详见示例
"timestamp":"",//是 	19 	发送请求的时间，格式"yyyy-MM-dd HH:mm:ss" 	2014-07-24 03:07:50
"version":"",//是 	3 	调用的接口版本，固定为：1.0 	1.0
"notify_url":"",//否 	256 	支付宝服务器主动通知商户服务器里指定的页面http/https路径。 	http://api.test.alipay.net/atinterface/receive_notify.htm
"app_auth_token":"",//否 	40 	详见应用授权概述
"biz_content":"",//	是 		请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
}
type AutoGenerated struct {
	AppID        string `json:"app_id"`
	Method       string `json:"method"`
	Format       string `json:"format"`
	ReturnURL    string `json:"return_url"`
	Charset      string `json:"charset"`
	SignType     string `json:"sign_type"`
	Sign         string `json:"sign"`
	Timestamp    string `json:"timestamp"`
	Version      string `json:"version"`
	NotifyURL    string `json:"notify_url"`
	AppAuthToken string `json:"app_auth_token"`
	BizContent   string `json:"biz_content"`
}
*/

package alipay

var AlipayERROR = map[string]string{
	"isv.grant-type-invalid":      "grant_type参数不正确",               //grant_type必须是authorization_code、refresh_token二者之一 若传入authorization_code为code换取令牌，若传入refresh_token为刷新令牌
	"isv.code-invalid":            "授权码(auth_code) <br>错误、状态不对或过期", //使用有效的auth_code重新执行令牌换取，或引导用户重新授权
	"isv.refresh-token-invalid":   "刷新令牌(refresh_token)错误或状态不对",    //使用有效的refresh_token重新执行令牌刷新，或引导用户重新授权
	"isv.refresh-token-time-out":  "刷新令牌(refresh_token)过期",         //使用有效的refresh_token重新执行令牌刷新，或引导用户重新授权
	"isv.refreshed-token-invalid": "刷新出来的令牌无效",                     //使用返回的刷新令牌再次刷新
	"isv.invalid-app-id":          "调用接口的应用标识(app_id)与令牌授权的应用不相符",  //传入正确的app_id和令牌，若开发者支付宝账号名下有多个app_id，或者开发者管理多个归属于不同支付宝账号的app_id，请注意不要混用不同app_id的code
	"isp.unknow-error":            "未知错误",                          //
}

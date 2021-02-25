package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/ooncn/common/Handler"
	"github.com/ooncn/common/util"
	"github.com/ooncn/opay/alipay"
)

type Controller struct {
	Handler.Context
}

var PublicKey = "支付宝公钥"

func (c Controller) getReturn(ctx iris.Context) {
	c.SetCxt(ctx)
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	fmt.Println(c.Params().Get("id"))
	m := c.URLParams()
	fmt.Println("params", util.JsonToStr(m))
	delete(m, "sign_type")
	err := alipay.AlipayVerifyMap(m, PublicKey)
	if err == nil {
		fmt.Println(m)
	} else {
		fmt.Println("验签失败", err)
	}
	c.RespSuccess()
}
func (c Controller) postReturn(ctx iris.Context) {
	c.SetCxt(ctx)
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	fmt.Println(c.Params().Get("id"))
	body := c.BodyStr()
	fmt.Println(body)
	m := util.URLParamsToMap(body)
	delete(m, "sign_type")
	i := alipay.AlipayVerifyMap(m, PublicKey)
	if i == nil {
		fmt.Println(m)
	} else {
		fmt.Println("验签失败", i)
	}
	fmt.Println("body", util.JsonToStr(m))
	c.RespSuccess()
}
func main() {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, Handler.NotFoundHandler)
	//app.Use(iris.Cache304(1*time.Minute))
	index := app.Party("/", Controller{}.Cors).AllowMethods(iris.MethodOptions)
	index.Get("/pay/notify/{id}", Controller{}.getReturn)
	index.Post("/pay/notify/{id}", Controller{}.postReturn)
	err := app.Run(iris.Addr(":90"),
		iris.WithoutPathCorrectionRedirection,
		iris.WithoutBodyConsumptionOnUnmarshal)
	if err != nil {
		app.Logger().Warn("Shutdown with error: " + err.Error())
	}
}

package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/ooncn/common/Handler"
	"github.com/ooncn/common/util"
)

type Controller struct {
	Handler.Context
}

func (c Controller) index(ctx iris.Context) {
	fmt.Println(c.URLParams())
	var i map[string]interface{}
	c.ReadJSON(&i)
	fmt.Println(util.JsonToStr(i))
	c.RespSuccess()
}
func main() {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, Handler.NotFoundHandler)
	//app.Use(iris.Cache304(1*time.Minute))
	index := app.Party("/", Controller{}.Cors).AllowMethods(iris.MethodOptions)
	index.Any("/pay/return", Controller{}.index)
	err := app.Run(iris.Addr(":90"),
		iris.WithoutPathCorrectionRedirection,
		iris.WithoutBodyConsumptionOnUnmarshal)
	if err != nil {
		app.Logger().Warn("Shutdown with error: " + err.Error())
	}
}

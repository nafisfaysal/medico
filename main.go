package main

import (
	models "charts/models"
	_ "charts/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/postgres"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterModel(new(models.AuthUser), new(models.Patient))
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres",
		"user=root password=41234 host=localhost")

	beego.BConfig.WebConfig.Session.SessionProvider = "postgresql"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = ""
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.CopyRequestBody = true
	beego.SetStaticPath("/static", "static")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/fonts", "static/fonts")
	beego.SetStaticPath("/testdata", "static/testdata")
	beego.SetViewsPath("views")
	beego.Run()
}

package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}


type ProtectedController struct {
	ProtectedApiController
}

func (c *ProtectedController) Login() {
	if c.GetUserID() > 0 {
		c.Redirect("/allcharts", http.StatusUnauthorized)
	}
}

func (c *ProtectedController) Patients() {
	if !c.IsAuthenticated() {
		c.Redirect("/login", http.StatusUnauthorized)
	} else {
		c.TplName = "patients.tpl"
	}
}

func (c *ProtectedController) Allcharts() {
	if !c.IsAuthenticated() {
		c.Redirect("/login", http.StatusUnauthorized)
	} else {
		c.TplName = "allcharts.tpl"
	}
}

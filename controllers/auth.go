package controllers

import (
	"encoding/json"

	"charts/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

const (
	UserIdSessionKey = "user_id"
)

type LoginController struct {
	ApiController
}

func (c *LoginController) Login() {
	c.TplName = "login.tpl"

	var req struct {
		Email    string `valid:"Required" json:"email"`
		Password string `valid:"Required" json:"password"`
	}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.SendError(err)
		return
	}

	valid := validation.Validation{}
	if ok, err := valid.Valid(&req); err != nil {
		c.SendError(err)
		return
	} else if !ok {
		c.SendValidationErrors(valid.Errors)
		return
	}

	var usr models.AuthUser

	o := orm.NewOrm()
	if err := o.QueryTable(&usr).Filter("UserEmail", req.Email).One(&usr); err != nil {
		if err == orm.ErrNoRows {
			c.SendError(models.ErrInvalidUserOrPassword)
			return
		} else {
			c.SendError(models.ErrUnknown)
			return
		}
	}

	if ok, err := usr.CheckPassword(req.Password); err != nil {
		c.SendError(models.ErrUnknown)
	} else if !ok {
		c.SendError(models.ErrInvalidUserOrPassword)
	} else {
		c.SetSession(UserIdSessionKey, usr.Id)
		c.SendEmptyResponse()
	}
}

type LogoutController struct {
	ApiController
}

func (c *LogoutController) Logout() {
	c.DelSession(UserIdSessionKey)
	c.SendEmptyResponse()
}

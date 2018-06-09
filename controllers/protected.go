package controllers

import (
	"net/http"

	"strconv"

	"errors"
	"fmt"
	"strings"

	"charts/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) SendError(err error) {
	if err == orm.ErrNoRows {
		err = models.ErrRecordNotFound
	}

	c.Data["json"] = map[string]interface{}{
		"error": err.Error(),
	}
	c.ServeJSON()
}

func (c *ApiController) SendValidationErrors(err []*validation.Error) {
	var arr []string
	for _, el := range err {
		arr = append(arr, fmt.Sprintf("%s: %s", el.Key, el.Message))
	}

	c.SendError(errors.New(strings.Join(arr, "\n")))
}

func (c *ApiController) SendResponse(res interface{}) {
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *ApiController) SendEmptyResponse() {
	c.Data["json"] = map[string]interface{}{}
	c.ServeJSON()
}

func (c *ApiController) GetID(param ...string) int {
	var err error
	var id int

	if len(param) > 0 && param[0] != "" {
		id, err = strconv.Atoi(c.Ctx.Input.Param(param[0]))
	} else {
		id, err = strconv.Atoi(c.Ctx.Input.Param(":id"))
	}

	if err != nil {
		return 0
	}

	return id
}

type ProtectedApiController struct {
	ApiController
}

func (c *ProtectedApiController) GetUserID() int {
	if ses := c.GetSession("user_id"); ses == nil {
		return 0
	} else {
		return ses.(int)
	}
}

func (c *ProtectedApiController) IsAuthenticated() bool {
	return c.GetUserID() != 0
}

func (c *ProtectedApiController) Prepare() {
	if !c.IsAuthenticated() {
		// not authorized
		c.Ctx.Output.Status = http.StatusUnauthorized
		c.SendError(models.ErrNotAuthorized)
	}
}

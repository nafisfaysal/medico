package controllers

import (
	"time"

	"encoding/json"

	"charts/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type PatientApiController struct {
	ProtectedApiController
}

func (c *PatientApiController) List() {
	var userId = c.GetUserID()
	var patients []*models.Patient

	o := orm.NewOrm()
	cond :=
		orm.NewCondition().
			And("Recycle", false).
			AndCond(orm.NewCondition().
				Or("Practice1", userId).
				Or("Practice2", userId).
				Or("Practice3", userId).
				Or("Practice4", userId).
				Or("Practice5", userId).
				Or("Practice6", userId).
				Or("Practice7", userId).
				Or("Practice8", userId).
				Or("Practice9", userId).
				Or("Practice10", userId))

	if _, err := o.QueryTable(new(models.Patient)).SetCond(cond).All(&patients); err != nil {
		c.SendError(err)
		return
	}

	var res []map[string]interface{}
	for _, p := range patients {
		res = append(res, patientToJson(p))
	}

	c.SendResponse(map[string]interface{}{
		"data": res,
	})
}

func (c *PatientApiController) Create() {
	var req PatientRec

	
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

	var p = models.Patient{
		Practice1:   c.GetUserID(),
		LastName:    req.LastName,
		FirstName:   req.FirstName,
		MiddleName:  req.LastName,
		DateOfBirth: req.DateOfBirth,
		CaseName:    req.CaseName,
	}

	o := orm.NewOrm()
	if _, err := o.Insert(&p); err != nil {
		c.SendError(err)
		return
	}

	if created, err := models.GetPatientById(c.GetUserID(), p.Id); err != nil {
		c.SendError(err)
		return
	} else {
		c.SendResponse(map[string]interface{}{
			"patient": patientToJson(created),
		})
		return
	}
}

func (c *PatientApiController) Get() {
	if p, err := models.GetPatientById(c.GetUserID(), c.GetID()); err != nil {
		c.SendError(err)
	} else {
		c.SendResponse(map[string]interface{}{
			"patient": patientToJson(p),
		})
	}
}

func (c *PatientApiController) Delete() {
	if p, err := models.GetPatientById(c.GetUserID(), c.GetID()); err != nil {
		c.SendError(err)
	} else {
		p.Recycle = true
		if _, err := orm.NewOrm().Update(p, "Recycle"); err != nil {
			c.SendError(err)
		} else {
			c.SendEmptyResponse()
		}
	}
}

func (c *PatientApiController) Update() {
	if p, err := models.GetPatientById(c.GetUserID(), c.GetID()); err != nil {
		c.SendError(err)
	} else {
		var req PatientRec

		// check incoming body is json and valid
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

		p.FirstName = req.FirstName
		p.LastName = req.LastName
		p.MiddleName = req.MiddleName
		p.DateOfBirth = req.DateOfBirth
		p.CaseName = req.CaseName

		if _, err := orm.NewOrm().Update(p, "FirstName", "LastName", "MiddleName", "DateOfBirth", "CaseName"); err != nil {
			c.SendError(err)
		} else {

			c.SendResponse(map[string]interface{}{
				"patient": patientToJson(p),
			})
		}
	}
}

type PatientRec struct {
	LastName    string    `json:"last_name" valid:"Required"`
	FirstName   string    `json:"first_name" valid:"Required"`
	MiddleName  string    `json:"middle_name"`
	DateOfBirth time.Time `json:"date_of_birth" valid:"Required"`
	CaseName    string    `json:"case_name"`
}

func patientToJson(p *models.Patient) map[string]interface{} {
	return map[string]interface{}{
		"id":            p.Id,
		"last_name":     p.LastName,
		"first_name":    p.FirstName,
		"middle_name":   p.MiddleName,
		"date_of_birth": p.DateOfBirth.Format(time.RFC3339),
		"case_name":     p.CaseName,
	}
}

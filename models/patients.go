package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Patient struct {
	Id          int
	Practice1   int       `orm:"column(practice1)"`
	Practice2   int       `orm:"column(practice2)"`
	Practice3   int       `orm:"column(practice3)"`
	Practice4   int       `orm:"column(practice4)"`
	Practice5   int       `orm:"column(practice5)"`
	Practice6   int       `orm:"column(practice6)"`
	Practice7   int       `orm:"column(practice7)"`
	Practice8   int       `orm:"column(practice8)"`
	Practice9   int       `orm:"column(practice9)"`
	Practice10  int       `orm:"column(practice10)"`
	LastName    string    `orm:"column(lastname)"`
	FirstName   string    `orm:"column(firstname)"`
	MiddleName  string    `orm:"column(middlename)"`
	DateOfBirth time.Time `orm:"column(dob)"`
	CaseName    string    `orm:"column(casename)"`
	Created     time.Time `orm:"column(created);type(datetime);auto_now_add"`
	Recycle     bool      `orm:"column(recycle)"`
}

func (Patient) TableName() string {
	return "patients"
}

func GetPatientById(userId int, id int) (*Patient, error) {
	p := new(Patient)

	cond := orm.NewCondition().
		And("Id", id).
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

	if err := orm.NewOrm().QueryTable(p).SetCond(cond).One(p); err != nil {
		return nil, err
	} else {
		return p, nil
	}
}

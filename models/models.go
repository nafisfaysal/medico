package models

import (
	"github.com/astaxie/beego/orm"
)

type Practices struct {
	Id               int
	PracticeName     string
	PracticeAddress1 string
	PracticeAddress2 string
	PracticeCity     string
	PracticeState    string
	PracticeZip      string
	PracticeCreated  string
}

func init() {
	orm.RegisterModel(new(Practices))
}

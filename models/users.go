package models

import (
	"crypto/md5"
	"encoding/hex"
)

type AuthUser struct {
	Id             int
	UserEmail      string `orm:"column(useremail);unique"`
	UserPassword   string `orm:"column(userpassword)"`
}

func (AuthUser) TableName() string {
	return "users"
}

func (u *AuthUser) CheckPassword(pass string) (bool, error) {
	hasher := md5.New()
	_, err := hasher.Write([]byte(pass))
	if err != nil {
		return false, err
	}

	str := hex.EncodeToString(hasher.Sum(nil))
	return str == u.UserPassword, nil
}

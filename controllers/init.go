package controllers

import (
	"github.com/astaxie/beego"
)

var secretToken string

// InitSecretToken from app config file
func InitSecretToken() {
	secretToken = beego.AppConfig.String("app.token")
}

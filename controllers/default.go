package controllers

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

// MainController is used to handle weChat requests
// the Post method handles weChat messages sent by users
type MainController struct {
	beego.Controller
	timestamp string
	nonce     string
	signature string
}

// Prepare is used to check the sha1sum
func (self *MainController) Prepare() {
	self.signature = self.GetString("signature")
	self.timestamp = self.GetString("timestamp")
	self.nonce = self.GetString("nonce")
	a := []string{self.nonce, self.timestamp, secretToken}
	sort.Strings(a)
	s := strings.Join(a, "")
	h := sha1.New()
	_, err := io.WriteString(h, s)
	if err != nil {
		beego.Error(err)
		self.Abort("500")
	}
	signature := fmt.Sprintf("%x", h.Sum(nil))
	if signature != self.signature {
		beego.Error("invalid signature")
		self.Abort("403")
	}
}

// Get method is used to authenticate the url
func (self *MainController) Get() {
	self.Ctx.WriteString(self.GetString("echostr"))
}

// Post method handles weChat messages sent by users
func (self *MainController) Post() {
	c, err := ioutil.ReadAll(self.Ctx.Request.Body)
	if err != nil {
		beego.Error("read request error")
		self.Abort("500")
	}
	msg := new(Message)
	err = xml.Unmarshal(c, msg)
	if err != nil {
		beego.Error(err)
		self.Abort("500")
	}
	returnMsg := msg.Handle()
	self.Data["xml"] = &returnMsg
	self.ServeXml()
}

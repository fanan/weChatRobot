package controllers

import (
	"encoding/xml"
)

type Message struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   ToUserName
	FromUserName FromUserName
	MsgType      MsgType
	Content      Content
	Event        Event
	EventKey     EventKey
	CreateTime   int
	MsgId        int
}

type ToUserName struct {
	Data string `xml:",chardata"`
}

type FromUserName struct {
	Data string `xml:",chardata"`
}

type Content struct {
	Data string `xml:",chardata"`
}

type MsgType struct {
	Data string `xml:",chardata"`
}

type Event struct {
	Data string `xml:",chardata"`
}

type EventKey struct {
	Data string `xml:",chardata"`
}

func (self *Message) Handle() (respMsg Message) {
	respMsg.FromUserName.Data = self.ToUserName.Data
	respMsg.ToUserName.Data = self.FromUserName.Data
	respMsg.MsgType.Data = TextType
	switch self.MsgType.Data {
	case TextType:
		respMsg.Content.Data = self.HandleText()
	case EventType:
		respMsg.Content.Data = self.HandleEvent()
	default:
		respMsg.Content.Data = "not supported yet"
	}
	return respMsg
}

func (self *Message) HandleText() string {
	return ""
}

func (self *Message) HandleEvent() string {
	return ""
}

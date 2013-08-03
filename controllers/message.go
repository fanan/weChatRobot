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

func HandleMessage(c []byte) (msg *Message) {
	return msg
}

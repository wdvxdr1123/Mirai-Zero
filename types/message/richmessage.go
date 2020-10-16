package message

import (
	"fmt"
	"github.com/Mrs4s/MiraiGo/message"
)

func NewRichMessage(r ...func(*RichMessage)) *RichMessage {
	m := &RichMessage{elems: []message.IMessageElement{}}
	return m.Append(r...)
}

func (m *RichMessage) Append(r ...func(*RichMessage)) *RichMessage {
	for _, f := range r {
		f(m)
	}
	return m
}

func Text(str ...interface{}) func(*RichMessage) {
	return func(richMessage *RichMessage) {
		richMessage.elems = append(richMessage.elems, message.NewText(fmt.Sprint(str...)))
	}
}

// 图片
func Image(fn ...func() ([]byte, error)) func(*RichMessage) {
	return func(richMessage *RichMessage) {
		for _, f := range fn {
			if v, err := f(); err == nil {
				richMessage.elems = append(richMessage.elems, message.NewImage(v))
			}
		}
	}
}

// QQ表情
func Face(faces ...int32) func(*RichMessage) {
	return func(richMessage *RichMessage) {
		for _, index := range faces {
			richMessage.elems = append(richMessage.elems, message.NewFace(index))
		}
	}
}

// At 消息
func At(uins ...int64) func(*RichMessage) {
	return func(richMessage *RichMessage) {
		for _, uin := range uins {
			richMessage.elems = append(richMessage.elems, message.NewAt(uin, "test"))
		}
	}
}

func AtAll() func(*RichMessage) {
	return func(richMessage *RichMessage) {
		richMessage.elems = append(richMessage.elems, message.AtAll())
	}
}

// todo
func Reply(file string) func(*RichMessage) {
	panic("impl me")
}

// todo
func LightApp(file string) func(*RichMessage) {
	panic("impl me")
}

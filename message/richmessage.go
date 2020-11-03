package message

import (
	"fmt"
	"github.com/Mrs4s/MiraiGo/message"
)

func NewMessageBuilder() *MessageBuilder {
	builder := &MessageBuilder{Elems: []message.IMessageElement{}}
	return builder
}

func NewMessageBuilderF(f func(builder *MessageBuilder)) *MessageBuilder {
	builder := &MessageBuilder{Elems: []message.IMessageElement{}}
	f(builder)
	return builder
}

func (builder *MessageBuilder) Text(str ...interface{}) *MessageBuilder {
	builder.Elems = append(builder.Elems, message.NewText(fmt.Sprint(str...)))
	return builder
}

// 图片
func (builder *MessageBuilder) Image(fn ...func() ([]byte, error)) *MessageBuilder {
	for _, f := range fn {
		if v, err := f(); err == nil {
			builder.Elems = append(builder.Elems, message.NewImage(v))
		}
	}
	return builder
}

// QQ表情
func (builder *MessageBuilder) Face(faces ...int32) *MessageBuilder {
	for _, index := range faces {
		builder.Elems = append(builder.Elems, message.NewFace(index))
	}
	return builder
}

// At 消息
func (builder *MessageBuilder) At(uins ...int64) *MessageBuilder {
	for _, uin := range uins {
		builder.Elems = append(builder.Elems, message.NewAt(uin))
	}
	return builder
}

func (builder *MessageBuilder) AtAll() *MessageBuilder {
	builder.Elems = append(builder.Elems, message.AtAll())
	return builder
}

// todo
func (builder *MessageBuilder) Reply(file string) *MessageBuilder {
	panic("impl me")
}

// todo
func (builder *MessageBuilder) LightApp(file string) *MessageBuilder {
	panic("impl me")
}

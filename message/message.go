package message

import "github.com/Mrs4s/MiraiGo/message"

// 只能单独发送的消息,群礼物，闪照，秀图之类的
type ISingleMessage interface {
	Send()
}

// 富文本消息，一般消息
type RichMessage struct {
	elems []message.IMessageElement
}

func (m *RichMessage) Append(r ...func(*RichMessage)) *RichMessage {
	for _, f := range r {
		f(m)
	}
	return m
}

// todo
func (m *RichMessage) Send() {
	panic("impl me")
}

package message

import "github.com/Mrs4s/MiraiGo/message"

type FlashPicMessage struct {
	message.ImageElement
}

func (g *FlashPicMessage) Send() {
	panic("impl me")
}

type ShowPicMessage struct {
	message.ImageElement
	effectId uint32
}

func (g *ShowPicMessage) Send() {
	panic("impl me")
}
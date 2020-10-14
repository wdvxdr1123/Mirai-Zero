package message

import (
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/wdvxdr1123/mirai-zero/zero"
)

// 能单独发送的消息,群礼物，闪照，秀图之类的
// 包括富文本消息
type IMessage interface {
	Send(zero.Accessory)
}

type (
	// 语音消息
	VoiceMessage struct {
		message.VoiceElement
	}

	// 富文本消息，一般消息
	RichMessage struct {
		elems []message.IMessageElement
	}

	// 群礼物
	GiftMessage struct {
		productId uint32
	}

	// 闪照
	FlashPicMessage struct {
		message.ImageElement
	}

	// 群秀图
	ShowPicMessage struct {
		message.ImageElement
		effectId uint32
	}
)

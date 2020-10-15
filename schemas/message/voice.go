package message

import (
	"github.com/wdvxdr1123/mirai-zero/zero"
)

// 语音消息初始化
func NewVoiceMessage(fn ...func() ([]byte, error)) *VoiceMessage {
	var voice = &VoiceMessage{}
	for _, f := range fn {
		if v, err := f(); err == nil {
			voice.Data = v
			break
		}
	}
	return voice
}

func (voice *VoiceMessage) Send(accessory zero.Accessory) {
	panic("impl me")
}

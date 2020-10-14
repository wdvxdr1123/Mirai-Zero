package message

import "github.com/Mrs4s/MiraiGo/message"

type VoiceMessage struct {
	message.VoiceElement
}

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

func (voice *VoiceMessage) Send() {
	panic("impl me")
}

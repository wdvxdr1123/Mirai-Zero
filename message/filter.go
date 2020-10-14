package message

import "github.com/Mrs4s/MiraiGo/message"

// 筛选指定的消息
func (r *RichMessage) filter(fn ...func(*message.IMessageElement) bool) *RichMessage { // todo:用pipeline改写
	var m = &RichMessage{}
	for _, elem := range r.elems {
		var st = true
		for _, f := range fn {
			st = st && f(&elem)
			if st != false {
				break
			}
		}
		if st {
			m.elems = append(m.elems, elem)
		}
	}
	return m
}

package message

import "github.com/Mrs4s/MiraiGo/message"

// 筛选指定的消息
func Filter(r *MessageBuilder, fn ...func(*message.IMessageElement) bool) *MessageBuilder { // todo:用pipeline改写
	var m = &MessageBuilder{}
	for _, elem := range r.Elems {
		var st = true
		for _, f := range fn {
			st = st && f(&elem)
			if st != false {
				break
			}
		}
		if st {
			m.Elems = append(m.Elems, elem)
		}
	}
	return m
}

func FirstMatched(r *MessageBuilder, fn ...func(element *message.IMessageElement) bool) *message.IMessageElement {
	for _, elem := range r.Elems {
		var st = true
		for _, f := range fn {
			st = st && f(&elem)
			if st != false {
				break
			}
		}
		if st {
			return &elem
		}
	}
	return nil
}

func Foreach(r *MessageBuilder, fn ...func(element message.IMessageElement) message.IMessageElement) *MessageBuilder {
	var m = &MessageBuilder{}
	for _, elem := range r.Elems {
		for _, f := range fn {
			m.Elems = append(m.Elems, f(elem))
		}
	}
	return m
}

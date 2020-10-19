package zero

import (
	"github.com/wdvxdr1123/mirai-zero/message"
	"github.com/wdvxdr1123/mirai-zero/types"
)

type MSG map[string]interface{}

type Handler func(z *Zero,event types.IEvent, session *Session)

type (
	ISession interface {
		Set(key string, value interface{})
		Get(key string) (interface{}, bool)
		Send()
		Reply()
		/* todo:目前看来比较难实现的方法
		Pause()
		Reject()
		Receive()
		Finish()
		*/
	}
)

type Session struct {
	// todo handler
	event   types.IEvent
	state   map[string]interface{}
	handler map[string]Handler
}

// 创建一个新会话
func NewSession(ms ...MSG) *Session {
	var s = &Session{}
	/*
	for _, m := range ms {
		for key, val := range m {
		}
	}
	 */
	return s
}

// 设置会话的属性
func (s *Session) Set(key string, val interface{}) {
	/*
	s.state.Store(key, val)
	// todo: 怎么把事件传过来 ？？？
	if handle, ok := s.handler.Load(key); ok {
		if f, ok := handle.(func(*Zero, *Session)); ok {
			go f(zero, s)
		}
	}
	 */
}

// 读取会话的属性
func (s *Session) Get(key string) (interface{}, bool) {
	panic("impl me")
}

// todo: 这部分返回值还没想好怎么弄
func (s *Session) Send(zero *Zero, message message.IMessage) (message.IMessage, error) {
	if e, ok := s.event.(types.IGroupEvent); ok {
		zero.SendGroupMessage(e.GroupId(), message)
	}
	if e, ok := s.event.(types.IPrivateEvent); ok { // todo
		panic(e)
		// zero.SendPrivateMessage(e.UserId(),message)
	}
	return nil, nil // 先空着
}

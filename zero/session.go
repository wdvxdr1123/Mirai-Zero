package zero

import (
	"github.com/wdvxdr1123/mirai-zero/message"
	"github.com/wdvxdr1123/mirai-zero/types"
	"sync"
)

type MSG map[string]interface{}

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
	state       sync.Map
}

// 创建一个新会话
func NewSession(ms ...MSG) *Session {
	var s = &Session{state: sync.Map{}}
	for _, m := range ms {
		for key, val := range m {
			s.state.Store(key, val)
		}
	}
	return s
}

// 设置会话的属性
func (s *Session) Set(key string, val interface{}) {
	s.state.Store(key, val)
	// todo: 属性改变事件(类似于nonebot2)
}

// 读取会话的属性
func (s *Session) Get(key string) (interface{}, bool) {
	return s.state.Load(key)
}

type BaseSession struct {
	Session
}

// todo: 这部分返回值还没想好怎么弄
func (s *BaseSession) Send(zero *Zero, event types.IMessageEvent, message message.IMessage) (message.IMessage, error) {
	if e, ok :=event.(types.IGroupEvent);ok {
		zero.SendGroupMessage(e.GroupId(),message)
	}
	if e, ok :=event.(types.IPrivateEvent);ok { // todo
		panic(e)
		// zero.SendPrivateMessage(e.UserId(),message)
	}
	return nil, nil // 先空着
}

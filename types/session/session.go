package session

import (
	"github.com/wdvxdr1123/mirai-zero/types"
	"github.com/wdvxdr1123/mirai-zero/types/message"
	"github.com/wdvxdr1123/mirai-zero/zero"
	"sync"
)

//go:generate stringer -type=Type
type Type uint

const (
	Base Type = 1 << iota
	Log

	Message
	Notice
	Request

	Group
	Private
)

type (
	ISession interface {
		Time() int64
		Type() Type
		Set(key string, value interface{})
		Get(key string) (interface{}, bool)
	}

	// 业务相关会话
	IBaseSession interface {
		// 发送消息
		Send(zero *zero.Zero, message *message.IMessage) (*message.IMessage, error)
		// 区分群聊和私聊
		SubType() Type
		// 群聊为群号，私聊与Sender的ID一致
		From() int64
		// 获取发送者信息
		Sender() types.IUser
	}

	// 消息会话
	IMessageSession interface {
		IBaseSession
		MessageType() Type
		// 原始消息
		Message() message.IMessage
		// 回复消息
		Reply(zero *zero.Zero, message *message.IMessage) (*message.IMessage, error)
		// 撤回消息
		Recall(zero *zero.Zero) error
	}

	// 通知会话
	INoticeSession interface {
		IBaseSession
		NoticeType() Type
	}

	// 请求会话
	IRequestSession interface {
		IBaseSession
		RequestType() Type
	}

	// 业务无关会话(客户端离线，日志之类的消息)
	ILogSession interface {
	}
)

type Session struct {
	state       sync.Map
	time        int32
	sessionType Type
}

// 创建一个新会话
func NewSession(tp Type, t int32, fn ...func() (string, interface{})) *Session {
	var s = &Session{sessionType: tp, time: t, state: sync.Map{}}
	for _, f := range fn {
		key, val := f()
		s.state.Store(key, val)
	}
	return s
}

// 获取事件发生的时间戳
func (s *Session) Time() int32 {
	return s.time
}

// 获取事件类型
func (s *Session) Type() Type {
	return s.sessionType
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

// 创建一个新会话
func NewBaseSession(tp Type, t int32, fn ...func() (string, interface{})) *BaseSession {
	var s = &BaseSession{Session{sessionType: tp, time: t, state: sync.Map{}}}
	// todo BaseSession
	for _, f := range fn {
		key, val := f()
		s.state.Store(key, val)
	}
	return s
}

// todo: 这部分返回值还没想好怎么弄
func (s *BaseSession) Send(zero *zero.Zero, message *message.IMessage) (*message.IMessage, error) {
	switch s.SubType() {
	case Group:
		zero.SendGroupMessage(s.From(), message)
	case Private:
		panic("impl me")
	default:
		panic("invalid session type")
	}
	return nil, nil // 先空着
}

func (s *BaseSession) From() int64 {
	var user interface{}
	if s.SubType() == Group {
		user, _ = s.state.Load("group_id")
	} else {
		user, _ = s.state.Load("user_id")
	}
	if sender, ok := user.(int64); ok {
		return sender
	}
	return 0
}

func (s *BaseSession) Sender() types.IUser {
	user, _ := s.state.Load("sender")
	if sender, ok := user.(types.IUser); ok {
		return sender
	}
	return nil
}

func (s *BaseSession) SubType() Type {
	return s.sessionType&Group + s.sessionType&Private
}

package zero

import (
	"github.com/wdvxdr1123/mirai-zero/message"
	"sync"
)

//go:generate stringer -type=SessionType
type SessionType uint

const (
	Base SessionType = 1 << iota
	Log

	Message
	Notice
	Request

	Group
	Private
)

type MSG map[string]interface{}

type (
	ISession interface {
		Time() int64
		Type() SessionType
		Set(key string, value interface{})
		Get(key string) (interface{}, bool)
	}

	// 业务相关会话
	IBaseSession interface {
		// 发送消息
		Send(zero *Zero, message *message.IMessage) (*message.IMessage, error)
		// 区分群聊和私聊
		SubType() SessionType
		// 群聊为群号，私聊与Sender的ID一致
		From() int64
		// 获取发送者信息
		Sender() MSG
	}

	// 消息会话
	IMessageSession interface {
		IBaseSession
		MessageType() SessionType
		// 原始消息
		Message() message.IMessage
		// 回复消息
		Reply(zero *Zero, message *message.IMessage) (*message.IMessage, error)
		// 撤回消息
		Recall(zero *Zero) error
	}

	// 通知会话
	INoticeSession interface {
		IBaseSession
		NoticeType() SessionType
	}

	// 请求会话
	IRequestSession interface {
		IBaseSession
		RequestType() SessionType
	}

	// 业务无关会话(客户端离线，日志之类的消息)
	ILogSession interface {
	}
)

type Session struct {
	state       sync.Map
	time        int32
	sessionType SessionType
}

// 创建一个新会话
func NewSession(tp SessionType, t int32, ms ...MSG) *Session {
	var s = &Session{sessionType: tp, time: t, state: sync.Map{}}
	for _, m := range ms {
		for key, val := range m {
			s.state.Store(key, val)
		}
	}
	return s
}

// 获取事件发生的时间戳
func (s *Session) Time() int32 {
	return s.time
}

// 获取事件类型
func (s *Session) Type() SessionType {
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
func NewBaseSession(tp SessionType, t int32,  ms ...MSG) *BaseSession {
	var s = &BaseSession{Session{sessionType: tp, time: t, state: sync.Map{}}}
	// todo BaseSession
	for _, m := range ms {
		for key, val := range m {
			s.state.Store(key, val)
		}
	}
	return s
}

// todo: 这部分返回值还没想好怎么弄
func (s *BaseSession) Send(zero *Zero, message message.IMessage) (message.IMessage, error) {
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
		user, _ = s.state.Load("group_code")
	} else {
		user, _ = s.state.Load("user_id")
	}
	if sender, ok := user.(int64); ok {
		return sender
	}
	return 0
}

func (s *BaseSession) Sender() MSG {
	user, _ := s.state.Load("sender")
	if sender, ok := user.(MSG); ok {
		return sender
	}
	return nil
}

func (s *BaseSession) SubType() SessionType {
	return s.sessionType&Group + s.sessionType&Private
}

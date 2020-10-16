package session

import (
	"github.com/wdvxdr1123/mirai-zero/types"
	"github.com/wdvxdr1123/mirai-zero/types/message"
	"github.com/wdvxdr1123/mirai-zero/zero"
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
	}

	// 业务相关会话
	IBaseSession interface {
		// 发送消息
		Send(zero *zero.Zero, message *message.IMessage) (*message.IMessage, error)
		// 区分群聊和私聊
		SubType()
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
	time        uint64
	sessionType Type
}

// 获取事件发生的时间戳
func (s *Session) Time() uint64 {
	return s.time
}

// 获取事件类型
func (s *Session) Type() Type {
	return s.sessionType
}

type BaseSession struct {
	Session
	from   int64
	sender types.IUser
}

// todo: 这部分返回值还没想好怎么弄
func (s *BaseSession) Send(zero *zero.Zero, message *message.IMessage) (*message.IMessage, error) {
	switch s.SubType() {
	case Group:
		zero.SendGroupMessage(s.from, message)
	case Private:
		panic("impl me")
	default:
		panic("invalid session type")
	}
	return nil, nil // 先空着
}

func (s *BaseSession) From() uint64 {
	return s.From()
}

func (s *BaseSession) Sender() types.IUser {
	return s.sender
}

func (s *BaseSession) SubType() Type {
	return s.sessionType&Group + s.sessionType&Private
}

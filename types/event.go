package types

import (
	"github.com/Mrs4s/MiraiGo/message"
)

type (
	IEvent interface {
		Time() int32
	}

	IGroupEvent interface {
		IEvent
		GroupId() int64
		UserId() int64
		Sender() User
	}

	IPrivateEvent interface {
		IEvent
		From() int64
		UserId() int64
		Sender() User
	}
)

type (
	IMessageEvent interface {
		MessageId() uint64
		RawMessage() []message.IMessageElement
	}
)

type (
	User struct {
		Uin      uint64 `json:"uin"`
		NickName string `json:"nick_name"`
		Sex      string `json:"sex"`
		Card     string `json:"card"`
		Role     string `json:"role"`
		Title    string `json:"title"`
	}

	BaseEvent struct {
		time int64
	}

	GroupEvent struct {
	}

	GroupMessageEvent struct {
		BaseEvent
		userId     int64
		groupId    int64
		sender     User
		rawMessage []message.IMessageElement
	}

	PrivateMessageEvent struct {
		BaseEvent
		userId     int64
		sender     User
		rawMessage []message.IMessageElement
	}
)

func (e *BaseEvent) Time() int64 {
	return e.time
}

func (e *GroupMessageEvent) Sender() User {
	return e.sender
}

func (e *GroupMessageEvent) GroupId() int64 {
	return e.groupId
}

func (e *GroupMessageEvent) UserId() int64 {
	return e.userId
}

func (e *GroupMessageEvent) RawMessage() []message.IMessageElement {
	return e.rawMessage
}

func (e *PrivateMessageEvent) Sender() User {
	return e.sender
}

func (e *PrivateMessageEvent) UserId() int64 {
	return e.userId
}

func (e *PrivateMessageEvent) RawMessage() []message.IMessageElement {
	return e.rawMessage
}

func NewBaseEvent(time int64) *BaseEvent {
	return &BaseEvent{
		time: time,
	}
}

func NewGroupMessageEvent(event *BaseEvent, userId, groupId int64, user User, msg ...message.IMessageElement) *GroupMessageEvent {
	return &GroupMessageEvent{
		BaseEvent:  *event,
		userId:     userId,
		groupId:    groupId,
		sender:     user,
		rawMessage: msg,
	}
}

func NewPrivateMessageEvent(event *BaseEvent, userId int64, user User, msg ...message.IMessageElement) *PrivateMessageEvent {
	return &PrivateMessageEvent{
		BaseEvent:  *event,
		userId:     userId,
		sender:     user,
		rawMessage: msg,
	}
}

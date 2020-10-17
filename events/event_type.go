package events

import (
	"github.com/wdvxdr1123/mirai-zero/types/session"
	"github.com/wdvxdr1123/mirai-zero/zero"
)

//go:generate stringer -type=ZeroEventsName
type ZeroEventsName uint8

const (
	ZeroEventGroupMessage ZeroEventsName = iota
	ZeroEventPrivateMessage
	ZeroEventGroupMute
	ZeroEventGroupRecalled
	ZeroEventFriendRecalled
	ZeroEventJoinGroup
	ZeroEventLeaveGroup
	ZeroEventMemberJoin
	ZeroEventMemberLeave
	ZeroEventMemberCardUpdated
	ZeroEventPermissionChanged
	ZeroEventJoinRequest
	ZeroEventFriendRequest
	ZeroEventNewFriend
	ZeroEventNotify
	ZeroEventGroupMessageReceipt

	ZeroEventLog
	ZeroEventDisconnect
	ZeroEventServerUpdated
)

type (
	EventGroupMessage        func(zero *zero.Zero, session session.IMessageSession)
	EventPrivateMessage      func(zero *zero.Zero, session session.IMessageSession)
	EventGroupMute           func(zero *zero.Zero, session session.IMessageSession)
	EventGroupRecalled       func(zero *zero.Zero, session session.IMessageSession)
	EventFriendRecalled      func(zero *zero.Zero, session session.IMessageSession)
	EventJoinGroup           func(zero *zero.Zero, session session.IMessageSession)
	EventLeaveGroup          func(zero *zero.Zero, session session.IMessageSession)
	EventMemberJoin          func(zero *zero.Zero, session session.IMessageSession)
	EventMemberLeave         func(zero *zero.Zero, session session.IMessageSession)
	EventMemberCardUpdated   func(zero *zero.Zero, session session.IMessageSession)
	EventPermissionChanged   func(zero *zero.Zero, session session.IMessageSession)
	EventJoinRequest         func(zero *zero.Zero, session session.IMessageSession)
	EventFriendRequest       func(zero *zero.Zero, session session.IMessageSession)
	EventNewFriend           func(zero *zero.Zero, session session.IMessageSession)
	EventNotify              func(zero *zero.Zero, session session.IMessageSession)
	EventGroupMessageReceipt func(zero *zero.Zero, session session.IMessageSession)

	EventLog           func(zero *zero.Zero, session session.ISession)
	EventDisconnect    func(zero *zero.Zero, session session.ISession)
	EventServerUpdated func(zero *zero.Zero, session session.ISession)
)

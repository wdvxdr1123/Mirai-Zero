package events

import (
	"context"
	"github.com/wdvxdr1123/mirai-zero/types"
	"github.com/wdvxdr1123/mirai-zero/types/message"
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
	EventGroupMessage        func(*context.Context, types.Group, message.IMessage)
	EventPrivateMessage      func(*context.Context, types.IUser, message.IMessage)
	EventGroupMute           func(*context.Context, types.Group, uint64)
	EventGroupRecalled       func(*context.Context)
	EventFriendRecalled      func(*context.Context)
	EventJoinGroup           func(*context.Context)
	EventLeaveGroup          func(*context.Context)
	EventMemberJoin          func(*context.Context)
	EventMemberLeave         func(*context.Context)
	EventMemberCardUpdated   func(*context.Context)
	EventPermissionChanged   func(*context.Context)
	EventJoinRequest         func(*context.Context)
	EventFriendRequest       func(*context.Context)
	EventNewFriend           func(*context.Context)
	EventNotify              func(*context.Context)
	EventGroupMessageReceipt func(*context.Context)

	EventLog           func(*context.Context)
	EventDisconnect    func(*context.Context)
	EventServerUpdated func(*context.Context)
)

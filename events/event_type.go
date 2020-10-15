package events

type ZeroEventsName uint8


//go:generate stringer -type=ZeroEventsName
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

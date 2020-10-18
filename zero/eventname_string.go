// Code generated by "stringer -type=EventName"; DO NOT EDIT.

package zero

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[GroupMessageEvent-0]
	_ = x[PrivateMessageEvent-1]
	_ = x[GroupMuteZeroEvent-2]
	_ = x[GroupRecalledEvent-3]
	_ = x[FriendRecalledEvent-4]
	_ = x[JoinGroupEvent-5]
	_ = x[LeaveGroupEvent-6]
	_ = x[MemberJoinEvent-7]
	_ = x[MemberLeaveEvent-8]
	_ = x[MemberCardUpdatedEvent-9]
	_ = x[PermissionChangedEvent-10]
	_ = x[JoinRequestEvent-11]
	_ = x[FriendRequestEvent-12]
	_ = x[NewFriendEvent-13]
	_ = x[NotifyEvent-14]
	_ = x[GroupMessageReceiptEvent-15]
	_ = x[LogEvent-16]
	_ = x[DisconnectEvent-17]
	_ = x[ServerUpdatedEvent-18]
}

const _EventName_name = "GroupMessageEventPrivateMessageEventGroupMuteZeroEventGroupRecalledEventFriendRecalledEventJoinGroupEventLeaveGroupEventMemberJoinEventMemberLeaveEventMemberCardUpdatedEventPermissionChangedEventJoinRequestEventFriendRequestEventNewFriendEventNotifyEventGroupMessageReceiptEventLogEventDisconnectEventServerUpdatedEvent"

var _EventName_index = [...]uint16{0, 17, 36, 54, 72, 91, 105, 120, 135, 151, 173, 195, 211, 229, 243, 254, 278, 286, 301, 319}

func (i EventName) String() string {
	if i >= EventName(len(_EventName_index)-1) {
		return "EventName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EventName_name[_EventName_index[i]:_EventName_index[i+1]]
}
package zero

//go:generate stringer -type=EventName
type EventName uint8

/*
const (
	GroupMessageEvent EventName = iota
	PrivateMessageEvent
	GroupMuteZeroEvent
	GroupRecalledEvent
	FriendRecalledEvent
	JoinGroupEvent
	LeaveGroupEvent
	MemberJoinEvent
	MemberLeaveEvent
	MemberCardUpdatedEvent
	PermissionChangedEvent
	JoinRequestEvent
	FriendRequestEvent
	NewFriendEvent
	NotifyEvent
	GroupMessageReceiptEvent

	LogEvent
	DisconnectEvent
	ServerUpdatedEvent
)

type (
	GroupMessageEventListener        func(zero *Zero, session IMessageSession)
	PrivateMessageEventListener      func(zero *Zero, session IMessageSession)
	GroupMuteEventListener           func(zero *Zero, session IMessageSession)
	GroupRecalledEventListener       func(zero *Zero, session IMessageSession)
	FriendRecalledEventListener      func(zero *Zero, session IMessageSession)
	JoinGroupEventListener           func(zero *Zero, session IMessageSession)
	LeaveGroupEventListener          func(zero *Zero, session IMessageSession)
	MemberJoinEventListener          func(zero *Zero, session IMessageSession)
	MemberLeaveEventListener         func(zero *Zero, session IMessageSession)
	MemberCardUpdatedEventListener   func(zero *Zero, session IMessageSession)
	PermissionChangedEventListener   func(zero *Zero, session IMessageSession)
	JoinRequestEventListener         func(zero *Zero, session IMessageSession)
	FriendRequestEventListener       func(zero *Zero, session IMessageSession)
	NewFriendEventListener           func(zero *Zero, session IMessageSession)
	NotifyEventListener              func(zero *Zero, session IMessageSession)
	GroupMessageReceiptEventListener func(zero *Zero, session IMessageSession)

	LogEventListener           func(zero *Zero, session ISession)
	DisconnectEventListener    func(zero *Zero, session ISession)
	ServerUpdatedEventListener func(zero *Zero, session ISession)
)
 */

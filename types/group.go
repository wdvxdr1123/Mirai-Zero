package types

type MemberPermission int

type Group struct {
	GroupCode   int64  `json:"group_code"`
	GroupName   string `json:"group_name"`
	MemberCount int32  `json:"member_count"`
}

type GroupMember struct {
	GroupCode int64            `json:"group_code"`
	UserId    int64            `json:"user_id"`
	NickName  string           `json:"nick_name"`
	Card      string           `json:"card"`
	Tile      string           `json:"tile"`
	Role      MemberPermission `json:"role"`
}

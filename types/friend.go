package types

type IUser interface {
	// 获取qq号
	Id() int64
	// 获取昵称
	NickName() string
	// todo: more information
}

type Friend struct {
	UserId   int64  `json:"user_id"`
	NickName string `json:"nick_name"`
	Remark   string `json:"remark"`
}

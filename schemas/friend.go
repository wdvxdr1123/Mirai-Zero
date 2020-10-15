package schemas

type Friend struct {
	UserId   int64  `json:"user_id"`
	NickName string `json:"nick_name"`
	Remark   string `json:"remark"`
}

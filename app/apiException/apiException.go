package apiException

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	InternalServerError     = NewError(200500, "系统异常，请稍后重试!")
	UsernameOrPasswordWrong = NewError(200501, "用户名或密码错误!")
	NoAccessPermission      = NewError(200502, "无访问权限!")
	HasBlocked              = NewError(200503, "拉黑关系已存在!")
	UsernameOrPasswordError = NewError(200505, "用户名或密码不符合规范!")
	ParamsError             = NewError(200506, "参数错误!")
	FileTypeError           = NewError(200506, "文件类型错误!")
	UsernameOccupied        = NewError(200507, "用户名已被占用!")
	NicknameOccupied        = NewError(200507, "昵称已被占用!")
	PostNotFound            = NewError(200508, "帖子不存在!")
	BlockNotFound           = NewError(200508, "拉黑不存在!")
	CommentNotFound         = NewError(200508, "评论不存在!")
	NoOperatePermission     = NewError(200509, "无操作权限!")
	AttemptToBlockSelf      = NewError(200510, "不能屏蔽自己的帖子!")
	FileTooLarge            = NewError(200511, "文件过大!")
	ContentTooLong          = NewError(200512, "内容过长!")
	NicknameTooLong         = NewError(200512, "昵称过长!")
)

func (e *Error) Error() string {
	return e.Msg
}

func NewError(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

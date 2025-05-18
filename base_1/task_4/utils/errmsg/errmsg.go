package errmsg

//错误处理常量及信息

const (
	SUCCESS      = 200
	ERROR        = 500
	ERROR_Params = 5001
	//用户名已被使用
	ERROR_USERNAME_USERD = 1001
	//用户名密码错误
	ERROR_PASSWORD_WRONG = 1002
	//用户不存在
	ERROR_USER_NOT_EXIST = 1003
	//TOKEN不存在
	ERROR_TOKEN_NOT_EXIST = 1004
	//TOKEN过期了
	ERROR_TOKEN_RUNTIEM = 1005
	//TOKEN错误
	ERROR_TOKEN_WRONG = 1006
	//TOKEN格式错误
	ERROR_TOKEN_TYPE_WRONG = 1007
	//用户无权限
	ERROR_USER_NO_RIGHT = 1008

	//code=2000...文章模块的错误
	ERROR_ARTICLE_NOT_EXIST   = 2001
	ERROR_ARTICLE_NO_COMMENTS = 2002
)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "Fail",
	ERROR_Params:           "参数有问题",
	ERROR_USERNAME_USERD:   "用户名已存在！",
	ERROR_PASSWORD_WRONG:   "密码错误！",
	ERROR_USER_NOT_EXIST:   "用户名不存在！",
	ERROR_TOKEN_NOT_EXIST:  "TOKEN不存在",
	ERROR_TOKEN_RUNTIEM:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
	ERROR_USER_NO_RIGHT:    "该用户无权限",

	ERROR_ARTICLE_NOT_EXIST:   "文章不存在!",
	ERROR_ARTICLE_NO_COMMENTS: "该文章没有评论",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}

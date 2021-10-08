package errcode

var (
	ErrorCreateUserFail    = NewError(4018001, "创建用户失败")
	ErrorUserExist         = NewError(4018002, "用户已存在")
	ErrorUserNotExist      = NewError(4018003, "用户未存在")
	ErrorPasswdWrong       = NewError(4018004, "密码错误")
	ErrorUpdateFail        = NewError(4018005, "更新用户失败")
	ErrorCreateProgramFail = NewError(5011001, "创建题目失败")
)

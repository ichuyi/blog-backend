package handler

const (
	ParaError         = 101
	ParaErrorMsg      = "参数错误"
	SQLError          = 102
	SQLErrorMsg       = "操作失败"
	UserNotExist      = 103
	UserNotExistMsg   = "用户不存在"
	PasswordError     = 104
	PasswordErrorMsg  = "密码错误"
	OpenFileError     = 105
	OpenFileErrorMsg  = "打开文件出错"
	ReadFileError     = 106
	ReadFileErrorMsg  = "读取文件出错"
	WriteFileError    = 107
	WriteFileErrorMsg = "写入文件出错"
	FileNotExist      = 108
	FileNotExistMsg   = "文件不存在"
	DeleteFileError=109
	DeleteFileErrorMsg="删除文件出错"
)

const (
	FINISH   = 1
	UNFINISH = 0
)

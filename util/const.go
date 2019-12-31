package util

const (
	ParaError          = 101
	ParaErrorMsg       = "参数错误"
	SQLError           = 102
	SQLErrorMsg        = "操作失败"
	UserNotExist       = 103
	UserNotExistMsg    = "用户不存在"
	PasswordError      = 104
	PasswordErrorMsg   = "密码错误"
	OpenFileError      = 105
	OpenFileErrorMsg   = "打开文件出错"
	ReadFileError      = 106
	ReadFileErrorMsg   = "读取文件出错"
	WriteFileError     = 107
	WriteFileErrorMsg  = "写入文件出错"
	FileNotExist       = 108
	FileNotExistMsg    = "文件不存在"
	DeleteFileError    = 109
	DeleteFileErrorMsg = "删除文件出错"
	OtherError         = 110
	OtherErrorMsg      = "其他错误"
	Unauthorized       = -99
	UnauthorizedMsg    = "未登录"
)

const (
	FINISH     = 1
	UNFINISH   = 0
	Secret     = "fmnwan"
	MaxAge     = 7 * 24 * 3600
	Domain     = "http://localhost"
	PrivateKey = `-----BEGIN RSA PRIVATE KEY-----
	MIICXgIBAAKBgQDexXA/QC3gMmFfyYaU3qPGugoKsquddr6uybef4SUIOrhcnBIw
WwN2XjkukiErr8vUFgZSmdQLlz3AGPZM+7UR4dxrZ+mZrgtGln1XpbxoILjeLU9F
/w5dQOnsQr+94UpQmoRiEGAdf/hC6EYDmbXOVwwqsW2vWCX8lKAK97i6sQIDAQAB
AoGBAIqL8pZz3NQ9oQ/IrxmxKdxzYcFrg44oBsmADOKzTKvEkVqPekR2pj2ctiV/
nn+kErlxhTckEpuu0SnCnJQeS2qNpZWt6eNpnTlQX1D7EAKWmoNi3Kzt9QsEY4+Q
XGhaj1+M99pf8/LAmYQplaWBjztw0Hc+xng87EvkM1Q26+YBAkEA8SXJP12atlvS
cIJgfgvRysPJAnwxtkAQk7fuwWIAutwh7iFCTgdNMwg42r/jVe8VGp45k06osLd2
+tZe8sDsoQJBAOx96MbdLGzJi3CvzpKAYlba2Qar9q3NOCkhj1cYrtrskHawnvdC
xRRdVvEnl8MqYAXhkKyUCzyPaNAHDt8mhBECQHqbVKQUCnpXWzp6/2Z2yfbG7qeF
z5yzG/qPFSRbmLVpARNa86RKkBS3RHjsAUEK6vb6pZlg7+HRlfvZDLNuJaECQQCI
fg8Yn9ShIR3itwVx1rlrSbpSuHOSUkykqKLzcOKSplCmwD+vlBDtNQYV/3T/BnkP
S+XVRUjK9jZXQouZzcARAkEA7b2FV2n1YbJa4tzlo8VckOabyJccZiUAaXZ5j5X3
c29PCT8EuZDcfAxR5r6nmcK15aKDCpj3g2SmEp6sk736Mw==
-----END RSA PRIVATE KEY-----`
)

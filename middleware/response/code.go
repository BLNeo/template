package response

// 错误码
const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	Unauthorized = 20001

	ErrorAuthCheckTokenTimeout = 20002
	ErrorAuthCheckTokenFail    = 20003
	//ErrorAuth                  = 20004
	//ErrorAuthNotExist          = 20005
	//
	//ErrorPhoneSignUpExist = 30001
)

var MsgFlags = map[int]string{
	SUCCESS:       "ok",
	ERROR:         "fail",
	InvalidParams: "请求参数错误",

	Unauthorized: "您需要先登录",

	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthCheckTokenFail:    "Token校验失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

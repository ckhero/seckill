package constant

const (
	ErrorOk = 0
	ErrorFail = 1
)

var errorMap = map[int]string {
	ErrorOk: "成功",
	ErrorFail: "失败",
}

func ErrorMsg(code int) string {
	if msg, ok := errorMap[code]; ok {
		return msg
	}
	return "";
}
package err_code

type codeMsg struct {
	Errno	int64
	Errmsg	string
}

var codeMap = map[string]codeMsg{
	"UNKNOWN_ERROR": {int64(11100), "UNKNOWN_ERROR" },
	"GET_USER_ERROR": {int64(11101), "GET_USER_ERROR" },
	"DELETE_USER_ERROR": {int64(11102), "DELETE_USER_ERROR" },
	"ADD_USER_ERROR": {int64(11102), "ADD_USER_ERROR" },
	"USERNAME_IS_EXIST_ERROR": {int64(11102), "USERNAME_IS_EXIST_ERROR" },
	"EMAIL_IS_EXIST_ERROR": {int64(11102), "EMAIL_IS_EXIST_ERROR" },
	"USER_IS_NOT_EXIST_ERROR": {int64(11102), "USER_IS_NOT_EXIST_ERROR" },
	"PASSWORD_ERROR": {int64(11102), "PASSWORD_ERROR" },
	"UPDATE_ERROR": {int64(11102), "UPDATE_ERROR" },
}

func Code(key string) (Errno int64, Errmsg string) {
	if codeMap[key].Errno == 0 {
		Errno = codeMap["UNKNOWN_ERROR"].Errno
		Errmsg = codeMap["UNKNOWN_ERROR"].Errmsg
		return
	}
	Errno = codeMap[key].Errno
	Errmsg = codeMap[key].Errmsg
	return
}

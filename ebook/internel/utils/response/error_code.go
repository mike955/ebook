package response

type codeMsg struct {
	Errno	int64
	Errmsg	string
}

var codeMap = map[string]codeMsg{
	"UNKNOWN_ERROR": {int64(11100), "UNKNOWN_ERROR" },
	
	// ------ ebook-user error 	11101 ~ 11199 ------ //
	"GET_USER_ERROR": {int64(11101), "GET_USER_ERROR" },
	"DELETE_USER_ERROR": {int64(11102), "DELETE_USER_ERROR" },
	"ADD_USER_ERROR": {int64(11102), "ADD_USER_ERROR" },
	"USERNAME_IS_EXIST_ERROR": {int64(11102), "USERNAME_IS_EXIST_ERROR" },
	"EMAIL_IS_EXIST_ERROR": {int64(11102), "EMAIL_IS_EXIST_ERROR" },
	"USER_IS_NOT_EXIST_ERROR": {int64(11102), "USER_IS_NOT_EXIST_ERROR" },
	"PASSWORD_ERROR": {int64(11102), "PASSWORD_ERROR" },
	"UPDATE_ERROR": {int64(11102), "UPDATE_ERROR" },
	
	
	// ------ ebook-privilege error	11200 ~ 11399 ------ //
	// role error	11200 ~ 11249
	"GET_ROLE_ERROR": {int64(11101), "GET_ROLE_ERROR" },
	"ROLE_IS_NOT_EXIST_ERROR": {int64(11101), "ROLE_IS_NOT_EXIST_ERROR" },
	"ROLE_IS_EXIST_ERROR": {int64(11101), "ROLE_IS_EXIST_ERROR" },
	"ADD_ROLE_ERROR": {int64(11101), "ADD_ROLE_ERROR" },
	
	// privilege error	11250 ~ 11299
	"GET_PRIVILEGE_ERROR": {int64(11101), "GET_PRIVILEGE_ERROR" },
	"PRIVILEGE_IS_EXIST_ERROR": {int64(11101), "PRIVILEGE_IS_EXIST_ERROR" },
	"ADD_PRIVILEGE_ERROR": {int64(11101), "ADD_PRIVILEGE_ERROR" },
	"PRIVILEGE_IS_NOT_EXIST_ERROR": {int64(11101), "PRIVILEGE_IS_NOT_EXIST_ERROR" },
	"DELETE_PRIVILEGE_ERROR": {int64(11101), "DELETE_PRIVILEGE_ERROR" },
	
	// user_role error 11299 ~ 11349
	"GET_USER_ROLE_ERROR": {int64(11101), "GET_USER_ROLE_ERROR" },
	"USER_ROLE_IS_EXIST_ERROR": {int64(11101), "USER_ROLE_IS_EXIST_ERROR" },
	"ADD_USER_ROLE_ERROR": {int64(11101), "ADD_USER_ROLE_ERROR" },
	"USER_ROLE_IS_NOT_EXIST_ERROR": {int64(11101), "USER_ROLE_IS_NOT_EXIST_ERROR" },
	"DELETE_USER_ROLE_ERROR": {int64(11101), "DELETE_USER_ROLE_ERROR" },
	"UPDATE_USER_ROLE_ERROR": {int64(11101), "UPDATE_USER_ROLE_ERROR" },
	
	// role_privilege error 11350 ~ 11399
	"GET_ROLE_PRIVILEGE_ERROR": {int64(11101), "GET_ROLE_PRIVILEGE_ERROR" },
	"ROLE_PRIVILEGE_IS_NOT_EXIST_ERROR": {int64(11101), "ROLE_PRIVILEGE_IS_NOT_EXIST_ERROR" },
	
	// ------ ebook-category error	11400 ~ 11499 ------ //
	
	// ------ ebook error	11500 ~ 11599 ------ //
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

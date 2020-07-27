package err_code

type codeMsg struct {
	Errno	int64
	Errmsg	string
}

var codeMap = map[string]codeMsg{
	"UNKNOWN_ERROR": {int64(11100), "UNKNOWN_ERROR" },

	// ebook-user error 	11101 ~ 11199
	"GET_USER_ERROR": {int64(11101), "GET_USER_ERROR" },
	"DELETE_USER_ERROR": {int64(11102), "DELETE_USER_ERROR" },
	"ADD_USER_ERROR": {int64(11103), "ADD_USER_ERROR" },
	"USERNAME_IS_EXIST_ERROR": {int64(11104), "USERNAME_IS_EXIST_ERROR" },
	"EMAIL_IS_EXIST_ERROR": {int64(11105), "EMAIL_IS_EXIST_ERROR" },
	"USER_IS_NOT_EXIST_ERROR": {int64(11106), "USER_IS_NOT_EXIST_ERROR" },
	"PASSWORD_ERROR": {int64(11107), "PASSWORD_ERROR" },
	"UPDATE_ERROR": {int64(11108), "UPDATE_ERROR" },
	

	// ebook-privilege error	11200 ~ 11399
			// role error	11200 ~ 11249
	"GET_ROLE_ERROR": {int64(11200), "GET_ROLE_ERROR" },
	"ROLE_IS_NOT_EXIST_ERROR": {int64(11201), "ROLE_IS_NOT_EXIST_ERROR" },
	"ROLE_IS_EXIST_ERROR": {int64(11202), "ROLE_IS_EXIST_ERROR" },
	"ADD_ROLE_ERROR": {int64(11203), "ADD_ROLE_ERROR" },

			// privilege error	11250 ~ 11299
	"GET_PRIVILEGE_ERROR": {int64(11250), "GET_PRIVILEGE_ERROR" },
	"PRIVILEGE_IS_EXIST_ERROR": {int64(11251), "PRIVILEGE_IS_EXIST_ERROR" },
	"ADD_PRIVILEGE_ERROR": {int64(11252), "ADD_PRIVILEGE_ERROR" },
	"PRIVILEGE_IS_NOT_EXIST_ERROR": {int64(11253), "PRIVILEGE_IS_NOT_EXIST_ERROR" },
	"DELETE_PRIVILEGE_ERROR": {int64(11254), "DELETE_PRIVILEGE_ERROR" },
	
			// user_role error 11300 ~ 11349
	"GET_USER_ROLE_ERROR": {int64(11300), "GET_USER_ROLE_ERROR" },
	"USER_ROLE_IS_EXIST_ERROR": {int64(11301), "USER_ROLE_IS_EXIST_ERROR" },
	"ADD_USER_ROLE_ERROR": {int64(11302), "ADD_USER_ROLE_ERROR" },
	"USER_ROLE_IS_NOT_EXIST_ERROR": {int64(11303), "USER_ROLE_IS_NOT_EXIST_ERROR" },
	"DELETE_USER_ROLE_ERROR": {int64(11304), "DELETE_USER_ROLE_ERROR" },
	"UPDATE_USER_ROLE_ERROR": {int64(11305), "UPDATE_USER_ROLE_ERROR" },

			// role_privilege error 11350 ~ 11399
	"GET_ROLE_PRIVILEGE_ERROR": {int64(11350), "GET_ROLE_PRIVILEGE_ERROR" },
	"ROLE_PRIVILEGE_IS_NOT_EXIST_ERROR": {int64(11351), "ROLE_PRIVILEGE_IS_NOT_EXIST_ERROR" },
	
	// ebook error	11400 ~ 11499
	"SIGNIN_ERROR": {int64(11400), "SIGNIN_ERROR" },
	
	
	// ebook-ebook error	11500 ~ 11599
	"ADD_EBOOK_ERROR": {int64(11500), "ADD_EBOOK_ERROR" },
	"GET_EBOOK_ERROR": {int64(11501), "GET_EBOOK_ERROR" },
	"EBOOK_IS_NOT_EXIST_ERROR": {int64(11502), "EBOOK_IS_NOT_EXIST_ERROR" },
	"DELETE_EBOOK_ERROR": {int64(11503), "DELETE_EBOOK_ERROR" },
	"UPDATE_EBOOK_ERROR": {int64(11504), "UPDATE_EBOOK_ERROR" },
	
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

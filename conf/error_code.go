package conf

type ERR_RESPONSE struct {
	errno int
	errmsg string
	data string
}

var ERROR_CODE_MAP = map[string]interface{}{
	"ACCOUNT_IS_EXIST_ERR" : &ERR_RESPONSE{
		errno:  4901,
		errmsg: "ACCOUNT_IS_EXIST",
		data:   "",
	},
	"SIGN_UP_ERR" : &ERR_RESPONSE{
		errno:  4902,
		errmsg: "SIGN_UP_ERR",
		data:   "",
	},
}
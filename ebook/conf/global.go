package conf

const (
		RUN_MODE = "debug"		// debug or release
		EBOOK_PORT = 8001
		READTIMEOUT = 600
		WRITETIMEOUT = 600
		
		
)

var (
	GRPC_ADDR_MAP = map[string]string{
		"user": "127.0.0.1:50801",
		"privilege": "127.0.0.1:50802",
	}
)
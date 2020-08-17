package conf

import "time"

const (
		RUN_MODE = "debug"		// debug or release
		EBOOK_PORT = 9401
		READTIMEOUT = 600
		WRITETIMEOUT = 600
		
		JWT_SECRET = "d56381bffae08ab3ee6297aedf3474e9"
		JWT_EXPIRE_TIME = 12 * 60 * 60 * time.Second
		
		EBOOK_UPLOAD_DIR = "/Users/superbear/dockerVolume/centos/nginx/ebook/static/file/"
		PREVIEW_UPLOAD_DIR = "/Users/superbear/dockerVolume/centos/nginx/ebook/static/image/"
)

var (
	GRPC_ADDR_MAP = map[string]string{
		"user": "127.0.0.1:50801",
		"privilege": "127.0.0.1:50802",
		"ebook": "127.0.0.1:50803",
	}
)
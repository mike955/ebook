package user

type SignUpRequest struct {
	Username     string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleId uint64 `json:"roleId" binding:"required"`
}


type SignInRequest struct {
	Username     string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignOutRequest struct {
	UserId     string `json:"userId" binding:"required"`
}

type GetRequest struct {
	UserId     string `json:"userId" binding:"required"`
}
package user

type SignUpRequest struct {
	Username     string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleId uint64 `json:"roleId" binding:"required"`
}

type SignUpResponse struct {
	error     string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInRequest struct {

}

type SignInResponse struct {

}

type SignOutRequest struct {

}

type SignOutResponse struct {

}
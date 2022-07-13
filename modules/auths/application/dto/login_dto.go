package dto

import "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/auths/domain/schema"

// LoginReqDTO type
type LoginReqDTO struct {
	schema.LoginRequest
}

// LoginResDTO type
type LoginResDTO struct {
	schema.LoginResponse
}

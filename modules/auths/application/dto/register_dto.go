package dto

import "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/auths/domain/schema"

// RegisterReqDTO type
type RegisterReqDTO struct {
	schema.RegisterRequest
}

// RegisterResDTO type
type RegisterResDTO struct {
	Email string `json:"email"`
}

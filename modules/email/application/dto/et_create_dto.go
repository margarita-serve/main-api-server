package dto

import (
	"encoding/json"

	domSchema "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/domain/schema/email_template"
)

// ETCreateReqDTO type
type ETCreateReqDTO struct {
	domSchema.ETCreateRequest
}

// ETCreateResDTO type
type ETCreateResDTO struct {
	domSchema.ETCreateResponse
}

// ToJSON covert to JSON
func (r *ETCreateResDTO) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}

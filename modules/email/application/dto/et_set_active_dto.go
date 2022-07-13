package dto

import (
	"encoding/json"

	domSchema "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/domain/schema/email_template"
)

// ETSetActiveReqDTO type
type ETSetActiveReqDTO struct {
	Keys *ETSetActiveKeysDTO `json:"keys"`
	Data *ETSetActiveDataDTO `json:"data"`
}

// ETSetActiveKeysDTO type
type ETSetActiveKeysDTO struct {
	domSchema.ETSetActiveKeys
}

// ETSetActiveDataDTO type
type ETSetActiveDataDTO struct {
	domSchema.ETSetActiveData
}

// ETSetActiveResDTO type
type ETSetActiveResDTO struct {
	domSchema.ETSetActiveResponse
}

// ToJSON covert to JSON
func (r *ETSetActiveResDTO) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}

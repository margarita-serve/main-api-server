package dto

import (
	"encoding/json"

	domSchema "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/domain/schema/email_template"
)

// ETFindByCodeReqDTO type
type ETFindByCodeReqDTO struct {
	domSchema.ETFindByCodeRequest
}

// ETFindByCodeResDTO type
type ETFindByCodeResDTO struct {
	Query domSchema.ETFindByCodeRequest `json:"query"`
	Data  domSchema.ETFindByCodeData    `json:"data"`
}

// ToJSON covert to JSON
func (r *ETFindByCodeResDTO) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}

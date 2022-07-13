package dto

import (
	"encoding/json"

	domSchema "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/domain/schema/email_template"
)

// ETListAllResDTO type
type ETListAllResDTO struct {
	domSchema.ETListAllResponse
}

// ToJSON covert to JSON
func (r *ETListAllResDTO) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}

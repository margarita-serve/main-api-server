package types

import "encoding/json"

type CreateServiceHealthRequest struct {
	InferenceName string `json:"inference_name"`
	ModelId       string `json:"model_id"`
}

func (r *CreateServiceHealthRequest) ToJSON() []byte {
	Mjson, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		return nil
	}
	return Mjson
}

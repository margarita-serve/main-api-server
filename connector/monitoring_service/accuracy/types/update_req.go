package types

import "encoding/json"

type UpdateAssociationIDRequest struct {
	InferenceName string `json:"inference_name"`
	AssociationID string `json:"association_id"`
}

type UpdateAssociationIDRequestDTO struct {
	AssociationID string `json:"association_id"`
}

func (r *UpdateAssociationIDRequestDTO) ToJSON() []byte {
	Mjson, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		return nil
	}
	return Mjson
}

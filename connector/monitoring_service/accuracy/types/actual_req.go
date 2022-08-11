package types

import "encoding/json"

type ActualRequest struct {
	InferenceName     string `json:"inference_name"`
	DatasetPath       string `json:"dataset_path"`
	ActualResponse    string `json:"actual_response"`
	AssociationColumn string `json:"association_column"`
}

func (r *ActualRequest) ToJSON() []byte {
	Mjson, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		return nil
	}
	return Mjson
}

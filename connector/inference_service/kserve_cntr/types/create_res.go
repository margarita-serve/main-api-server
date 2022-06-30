package types

type CreateInferenceServiceResponse struct {
	Message       string `json:"projectID"`
	Inferencename string `json:"modelPackageID"`
	Revision      string `json:"revision"`
	Url           string `json:"url"`
	Data          string `json:"data"`
}

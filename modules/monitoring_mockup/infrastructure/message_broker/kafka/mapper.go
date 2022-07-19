package kafka

type KafkaMsg struct {
	InferenceName string `json:"inference_name"`
	DriftResult   string `json:"drift_result"`
}

type OrgMsg struct {
	Msg     []byte
	MsgType string
}

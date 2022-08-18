package kafka

type KafkaMsg struct {
	InferenceName string `json:"inference_name"`
	Result        string `json:"result"`
}

type OrgMsg struct {
	Msg     []byte
	MsgType string
}

package entity

type WebHookEvent struct {
	WebHookID    string
	ID           string
	URL          string
	Method       string
	CustomHeader string
	MessageBody  string
	SendSuccess  bool
}

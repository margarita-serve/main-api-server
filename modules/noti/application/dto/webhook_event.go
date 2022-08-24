package dto

type WebHookEvent struct {
	ID            string
	Name          string
	TriggerSource string //Datadrift, Accuracy
	URL           string
	Method        string
	CustomHeader  string
	MessageBody   string
}

type GetWebHookEventRequestDTO struct {
	DeploymentID   string
	WebHookEventID string
}

type GetWebHookEventResponseDTO struct {
	WebHookEvent
}

type GetWebHookEventListRequestDTO struct {
	DeploymentID string
	Name         string `json:"name" extensions:"x-order=1"`                  // 검색조건: 배포 명
	Limit        int    `json:"limit" extensions:"x-order=2"`                 // 한번에 조회 할 건수
	Page         int    `json:"page" extensions:"x-order=3"`                  // 조회 할 페이지, 첫 조회후 TotalPages 범위 내에서 선택 후 보낸다
	Sort         string `enums:"CreateAsc,CreateDesc" extensions:"x-order=4"` //정열방식, CreateAsc: 생성시간 내림차순, CraeteDesc: 생성시간 역차순
}

type GetWebHookEventListResponseDTO struct {
	Limit      int
	Page       int
	Sort       string
	TotalRows  int64
	TotalPages int
	Rows       interface{}
}

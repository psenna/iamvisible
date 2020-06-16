package models

// CheckRequest A request to check the visibility
type CheckRequest struct {
	URL             string                 `json:"url"`
	Form            map[string]interface{} `json:"form"`
	QueryParameters map[string]string      `json:"query_parameters"`
	ExpectedStatus  int                    `json:"expected_status"`
	ReplyStatusOk   int                    `json:"reply_status_ok"`
	ReplyStatusNok  int                    `json:"reply_status_nok"`
	TimeOut         int                    `json:"time_out"`
}

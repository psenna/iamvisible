package models

import (
	"strconv"

	"github.com/psenna/isup-http-client/isuphttp"
)

// CheckRequest A request to check the visibility
type CheckRequest struct {
	URL             string                 `json:"url"`
	Method          string                 `json:"method"`
	Form            map[string]interface{} `json:"form"`
	Headers         map[string]interface{} `json:"headers"`
	QueryParameters map[string]interface{} `json:"query_parameters"`
	Insecure        bool                   `json:"insecure"`
	ExpectedStatus  int                    `json:"expected_status"`
	ReplyStatusOk   int                    `json:"reply_status_ok"`
	ReplyStatusNok  int                    `json:"reply_status_nok"`
	TimeOut         int                    `json:"timeout"`
}

// GetCheckRequest Instantiate a check request with default values
func GetCheckRequest() CheckRequest {
	return CheckRequest{
		URL:            "https://iamvisible.tk",
		Method:         "GET",
		ExpectedStatus: 200,
		ReplyStatusOk:  200,
		ReplyStatusNok: 500,
		TimeOut:        10000,
	}
}

// RunCheckRequest Peform the check request
func (cr CheckRequest) RunCheckRequest() (status int, message map[string]interface{}) {
	request := cr.getHTTPRequest()

	client := isuphttp.HTTPClient{}

	response := client.HTTPCall(request)

	if message == nil {
		message = make(map[string]interface{})
	}

	if response.StatusCode == cr.ExpectedStatus {
		status = cr.ReplyStatusOk
		message["Test"] = "OK"

	} else {
		status = cr.ReplyStatusNok
		message["Test"] = "NOK"
	}

	message["status"] = response.StatusCode
	message["error"] = response.Error
	message["response_time"] = strconv.FormatFloat(response.ResponseTime, 'f', 6, 64)

	return
}

func (cr CheckRequest) getHTTPRequest() isuphttp.HTTPRequest {
	httpRequest := isuphttp.GetHTTPRequest(cr.Method, cr.URL)

	httpRequest = httpRequest.SetInsecureRequest(cr.Insecure)

	if len(cr.Form) != 0 {
		httpRequest = httpRequest.SetFormParams(cr.Form)
	}

	if len(cr.QueryParameters) != 0 {
		httpRequest = httpRequest.SetQueryParams(cr.QueryParameters)
	}

	if len(cr.Headers) != 0 {
		httpRequest = httpRequest.SetHeaders(cr.Headers)
	}

	httpRequest = httpRequest.SetTimeOut(cr.TimeOut)

	return httpRequest
}

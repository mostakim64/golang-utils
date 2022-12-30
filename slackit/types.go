package slackit

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
)

type ClientRequest struct {
	Header      string   `json:"header"`
	ServiceName string   `json:"service_name"`
	Summary     string   `json:"summary"`
	Metadata    string   `json:"metadata"`
	Details     string   `json:"details"`
	Status      int      `json:"status"`
	Mentions    []string `json:"mentions"`
}

func (req *ClientRequest) Validate() error {
	return v.ValidateStruct(req,
		v.Field(&req.ServiceName, v.Required),
		v.Field(&req.Summary, v.Required),
		v.Field(&req.Details, v.Required),
		v.Field(&req.Status, v.Required),
	)
}

type ApiError struct {
	Api  string `json:"api"`
	Url  string `json:"url"`
	Data Data   `json:"data"`
}

type Data struct {
	Status       int         `json:"status"`
	Headers      interface{} `json:"headers"`
	RequestBody  interface{} `json:"request_body"`
	ResponseBody interface{} `json:"response_body"`
}

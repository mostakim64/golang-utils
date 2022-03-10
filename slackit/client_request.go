package slackit

type ClientRequest struct {
	ServiceName string `json:"service_name"`
	Summary     string `json:"summary"`
	Metadata    string `json:"metadata"`
	Details     string `json:"details"`
	Status      int `json:"status"`
}

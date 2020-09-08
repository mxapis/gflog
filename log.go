package gflog

type log struct {
	Message  interface{} `json:"message"`
	Severity Severity    `json:"severity"`
}

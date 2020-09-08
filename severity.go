package gflog

type Severity string

func (s Severity) String() string {
	return string(s)
}

const (
	Default   Severity = "default"
	Debug     Severity = "debug"
	Info      Severity = "info"
	Notice    Severity = "notice"
	Warning   Severity = "warning"
	Error     Severity = "error"
	Critical  Severity = "critical"
	Alert     Severity = "alert"
	Emergency Severity = "emergency"
)

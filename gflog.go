package gflog

import (
	"encoding/json"
	"fmt"
	"os"
)

func Print(s Severity, m ...interface{}) {
	log := &log{
		Message:  fmt.Sprint(m...),
		Severity: s,
	}

	bb, _ := json.Marshal(log)

	fmt.Printf("`%s`\n", bb)
}

func Printf(s Severity, format string, m ...interface{}) {
	log := &log{
		Message:  fmt.Sprintf(format, m...),
		Severity: s,
	}

	bb, _ := json.Marshal(log)

	fmt.Printf("`%s`\n", bb)
}

func Fatal(m ...interface{}) {
	Print(Emergency, m...)
	os.Exit(1)
}
func Fatalf(format string, m ...interface{}) {
	Printf(Emergency, format, m...)
	os.Exit(1)
}

package logger

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// list of popular formats (I'm dot sure if it's complete)
	formats = []string{
		"%v",
		"%+v",
		"%#v",
		"%T",
		"%t",
		"%d",
		"%b",
		"%c",
		"%x",
		"%f",
		"%e",
		"%E",
		"%s",
		"%q",
		"%x",
		"%p",
		"%6d",
		"%6.2f",
		"%-6.2f",
		"%6s",
		"%-6s",
	}
)

// formatter escape logs to prevent log injection
// by wrapping formats with `"` and escaping `\n` and `\r`
// and making error object ready to use
func formatter(format string, args ...any) error {
	for _, f := range formats {
		format = strings.ReplaceAll(format, f, "\""+f+"\"")
	}
	msg := fmt.Sprintf(format, args...)
	msg = strings.ReplaceAll(msg, "\n", " [ESCAPED] ")
	msg = strings.ReplaceAll(msg, "\r", " [ESCAPED] ")
	return errors.New(msg)
}

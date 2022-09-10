package logger

import "strings"

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

// escaper escape logs to prevent log injection
// by wrapping formats with `"` and escaping `\n` and `\r`
func escaper(format string) string {
	format = strings.ReplaceAll(format, "\n", " [ESCAPED] ")
	format = strings.ReplaceAll(format, "\r", " [ESCAPED] ")
	for _, f := range formats {
		format = strings.ReplaceAll(format, f, "\""+f+"\"")
	}
	return format
}

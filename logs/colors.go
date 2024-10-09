package logs

import (
	"fmt"
	"strconv"
)

// // ANSI color codes
// const (
//     Reset  = "\033[0m"
//     Red    = "\033[31m"
//     Green  = "\033[32m"
//     Yellow = "\033[33m"
//     Blue   = "\033[34m"
//     Purple = "\033[35m"
//     Cyan   = "\033[36m"
//     Gray   = "\033[90m"
// )

const (
	reset = "\033[0m"

	black        = 30
	red          = 31
	green        = 32
	yellow       = 33
	blue         = 34
	magenta      = 35
	cyan         = 36
	lightGray    = 37
	darkGray     = 90
	lightRed     = 91
	lightGreen   = 92
	lightYellow  = 93
	lightBlue    = 94
	lightMagenta = 95
	lightCyan    = 96
	white        = 97
)

func colorize(colorCode int, v string) string {
	return fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(colorCode), v, reset)
}

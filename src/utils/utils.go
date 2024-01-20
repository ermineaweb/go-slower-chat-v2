package utils

import "fmt"

var colors = map[string]string{
	"black":   "\x1b[30m",
	"red":     "\x1b[31m",
	"green":   "\x1b[32m",
	"yellow":  "\x1b[33m",
	"blue":    "\x1b[34m",
	"magenta": "\x1b[35m",
	"cyan":    "\x1b[36m",
	"white":   "\x1b[37m",
	"bold":    "\x1b[1m",
	"reset":   "\x1b[0m\x1b[39m",
}

func Format(text, color string) string {
	return fmt.Sprintf("%v%v%v", colors[color], text, colors["reset"])
}

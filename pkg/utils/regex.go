package utils

import (
	"fmt"
	"regexp"
)

const (
	number  = `\d*(?:\.\d*)?`
	percent = `(?:` + number + `%|0)`
	hue     = `(?:` + number + `)(?:deg|grad|rad|turn)?`
)

var Hex = regexp.MustCompile(`^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`)

var Rgb = regexp.MustCompile(fmt.Sprintf(
	`^rgb\(\s*(%s)\s*,\s*(%s)\s*,\s*(%s)\s*\)$`,
	number, number, number,
))

var Hsl = regexp.MustCompile(fmt.Sprintf(
	`^hsl\(\s*(%s)\s*,\s*(%s)\s*,\s*(%s)\s*\)$`,
	hue, percent, percent,
))

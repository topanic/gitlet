package utils

import (
	"fmt"
	"strconv"
)

// Color defines a single SGR Code
type Color int

// Foreground text colors
const (
    FgBlack Color = iota + 30
    FgRed
    FgGreen
    FgYellow
    FgBlue
    FgMagenta
    FgCyan
    FgWhite
)

// Foreground Hi-Intensity text colors
const (
    FgHiBlack Color = iota + 90
    FgHiRed
    FgHiGreen
    FgHiYellow
    FgHiBlue
    FgHiMagenta
    FgHiCyan
    FgHiWhite
)

// Colorize a string based on given color.
func Colorize(s string, c Color) string {
    return fmt.Sprintf("\033[1;%s;48m%s\033[0m", strconv.Itoa(int(c)), s)
}

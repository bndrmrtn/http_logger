package http_logger

import (
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/term"
	"strings"
	"time"
)

type Logger struct{}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Log(m string, p string) {
	var method string

	switch m {
	case "GET":
		method = color.New(color.FgHiGreen).Sprint(m)
		break
	case "POST":
		method = color.New(color.FgHiBlue).Sprint(m)
		break
	case "PUT":
		method = color.New(color.FgHiCyan).Sprint(m)
		break
	case "PATCH":
		method = color.New(color.FgHiYellow).Sprint(m)
		break
	case "DELETE":
		method = color.New(color.FgHiRed).Sprint(m)
		break
	default:
		method = color.New(color.FgHiMagenta).Sprint(m)
		break
	}

	t := time.Now().Format(time.TimeOnly)
	tColor := color.New(color.FgHiBlack).Sprint(t)

	l.print(len(fmt.Sprintf("%v %v %v", m, p, t)), fmt.Sprintf("%v %v", method, p), tColor)
}

func (l *Logger) Auth(mode string, user string) {
	l.print(8+len(mode)+len(user), "  🔒 "+color.New(color.FgYellow).Sprint(mode), user+"  ")
}

func (*Logger) print(cLen int, right string, left string) {
	w, _, err := term.GetSize(0)
	length := w - 1 - (cLen + 4)

	if err != nil || length <= 2 {
		fmt.Printf("  %v %v %v  \n", right, color.New(color.FgHiBlack).Sprint(".........."), left)
	} else {
		dots := strings.Repeat(".", length)
		fmt.Printf("  %v %v %v  \n", right, color.New(color.FgHiBlack).Sprint(dots), left)
	}
}

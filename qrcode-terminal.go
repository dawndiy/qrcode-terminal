package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/skip2/go-qrcode"
)

const (
	NormalBlack   = "\033[38;5;0m  \033[0m"
	NormalRed     = "\033[38;5;1m  \033[0m"
	NormalGreen   = "\033[38;5;2m  \033[0m"
	NormalYellow  = "\033[38;5;3m  \033[0m"
	NormalBlue    = "\033[38;5;4m  \033[0m"
	NormalMagenta = "\033[38;5;5m  \033[0m"
	NormalCyan    = "\033[38;5;6m  \033[0m"
	NormalWhite   = "\033[38;5;7m  \033[0m"

	BrightBlack   = "\033[48;5;0m  \033[0m"
	BrightRed     = "\033[48;5;1m  \033[0m"
	BrightGreen   = "\033[48;5;2m  \033[0m"
	BrightYellow  = "\033[48;5;3m  \033[0m"
	BrightBlue    = "\033[48;5;4m  \033[0m"
	BrightMagenta = "\033[48;5;5m  \033[0m"
	BrightCyan    = "\033[48;5;6m  \033[0m"
	BrightWhite   = "\033[48;5;7m  \033[0m"
)

var (
	frontColor      string
	backgroundColor string
	levelString     string
)

func init() {
	flag.Usage = printHelp
	flag.StringVar(&frontColor, "f", "black", "Front color")
	flag.StringVar(&backgroundColor, "b", "white", "Background color")
	flag.StringVar(&levelString, "l", "m", "Error correction level")
}

func main() {

	flag.Parse()

	var front, back, content string
	var err error
	var level qrcode.RecoveryLevel

	if front, err = parseColor(frontColor); err != nil {
		fmt.Println(err)
		printHelp()
		return
	}

	if back, err = parseColor(backgroundColor); err != nil {
		fmt.Println(err)
		printHelp()
		return
	}

	if level, err = parseLevel(levelString); err != nil {
		fmt.Println(err)
		printHelp()
		return
	}

	if content = flag.Arg(0); content == "" {
		content = "https://github.com/dawndiy/qrcode-terminal"
	}

	qr, err := qrcode.New(content, level)
	if err != nil {
		fmt.Println(err)
		return
	}

	for ir, row := range qr.Bitmap() {
		lr := len(row)
		if ir == 0 || ir == 1 || ir == 2 ||
			ir == lr-1 || ir == lr-2 || ir == lr-3 {
			continue
		}
		for ic, col := range row {
			lc := len(qr.Bitmap())
			if ic == 0 || ic == 1 || ic == 2 ||
				ic == lc-1 || ic == lc-2 || ic == lc-3 {
				continue
			}
			if col {
				fmt.Print(front)
			} else {
				fmt.Print(back)
			}
		}
		fmt.Println()
	}
}

func parseColor(str string) (color string, err error) {
	s := strings.ToUpper(str)
	switch s {
	case "BLACK":
		color = BrightBlack
	case "RED":
		color = BrightRed
	case "GREEN":
		color = BrightGreen
	case "YELLOW":
		color = BrightYellow
	case "BLUE":
		color = BrightBlue
	case "MAGENTA":
		color = BrightMagenta
	case "CYAN":
		color = BrightCyan
	case "WHITE":
		color = BrightWhite
	default:
		err = errors.New(fmt.Sprintf("'%s' is not support.", str))
	}

	return
}

func parseLevel(str string) (level qrcode.RecoveryLevel, err error) {
	s := strings.ToUpper(str)
	switch s {
	case "L":
		level = qrcode.Low
	case "M":
		level = qrcode.Medium
	case "Q":
		level = qrcode.High
	case "H":
		level = qrcode.Highest
	default:
		err = errors.New(fmt.Sprintf("'%s' is not support.", str))
	}

	return
}

func printHelp() {

	helpStr := `QRCode generater terminal edition.

Supported background colors: [black, red, green, yellow, blue, magenta, cyan, white]
Supported front colors: [black, red, green, yellow, blue, magenta, cyan, white]
Supported error correction levels: [L, M, Q, H]
`
	fmt.Println(helpStr)
	flag.PrintDefaults()
}

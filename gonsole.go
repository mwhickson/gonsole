package gonsole

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type ConsoleHandler struct {
	reader *bufio.Reader
}

func NewConsoleHandler() *ConsoleHandler {
	ch := &ConsoleHandler{reader: bufio.NewReader(os.Stdin)}
	return ch
}

func (ch *ConsoleHandler) EmitSequence(codes ...string) {
	sequence := "\x1b[" + strings.Join(codes, ";")
	fmt.Print(sequence)
}

func (ch *ConsoleHandler) Reset() {
	ch.EmitSequence("0m")
}

func (ch *ConsoleHandler) CursorToHome() {
	ch.EmitSequence("H")
}

func (ch *ConsoleHandler) HideCursor() {
	ch.EmitSequence("?25l")	
}

func (ch *ConsoleHandler) ShowCursor() {
	ch.EmitSequence("?25h")	
}

func (ch *ConsoleHandler) ClearScreen() {
	ch.EmitSequence("2J")
}

func (ch *ConsoleHandler) SetColor(color string) {
	colorSequence := color + "m" 
	ch.EmitSequence(colorSequence)	
}

func (ch *ConsoleHandler) SetForegroundBlack() {
	ch.SetColor("30")
}

func (ch *ConsoleHandler) SetForegroundRed() {
	ch.SetColor("31")
}

func (ch *ConsoleHandler) SetForegroundGreen() {
	ch.SetColor("32")
}

func (ch *ConsoleHandler) SetForegroundYellow() {
	ch.SetColor("33")
}

func (ch *ConsoleHandler) SetForegroundBlue() {
	ch.SetColor("34")
}

func (ch *ConsoleHandler) SetForegroundMagenta() {
	ch.SetColor("35")
}

func (ch *ConsoleHandler) SetForegroundCyan() {
	ch.SetColor("36")
}

func (ch *ConsoleHandler) SetForegroundWhite() {
	ch.SetColor("37")
}

func (ch *ConsoleHandler) SetForegroundDefault() {
	ch.SetColor("38")
}

func (ch *ConsoleHandler) SetBackgroundBlack() {
	ch.SetColor("40")
}

func (ch *ConsoleHandler) SetBackgroundRed() {
	ch.SetColor("41")
}

func (ch *ConsoleHandler) SetBackgroundGreen() {
	ch.SetColor("42")
}

func (ch *ConsoleHandler) SetBackgroundYellow() {
	ch.SetColor("43")
}

func (ch *ConsoleHandler) SetBackgroundBlue() {
	ch.SetColor("44")
}

func (ch *ConsoleHandler) SetBackgroundMagenta() {
	ch.SetColor("45")
}

func (ch *ConsoleHandler) SetBackgroundCyan() {
	ch.SetColor("46")
}

func (ch *ConsoleHandler) SetBackgroundWhite() {
	ch.SetColor("47")
}

func (ch *ConsoleHandler) SetBackgroundDefault() {
	ch.SetColor("48")
}

func (ch *ConsoleHandler) SetBold() {
	ch.EmitSequence("1m")
}

func (ch *ConsoleHandler) UnsetBold() {
	ch.EmitSequence("22m")
}

func (ch *ConsoleHandler) SetItalic() {
	ch.EmitSequence("3m")
}

func (ch *ConsoleHandler) UnsetItalic() {
	ch.EmitSequence("23m")
}

func (ch *ConsoleHandler) SetUnderline() {
	ch.EmitSequence("4m")
}

func (ch *ConsoleHandler) UnsetUnderline() {
	ch.EmitSequence("24m")
}

func (ch *ConsoleHandler) SetBlink() {
	ch.EmitSequence("5m")
}

func (ch *ConsoleHandler) UnsetBlink() {
	ch.EmitSequence("25m")
}

func (ch *ConsoleHandler) SetInverse() {
	ch.EmitSequence("7m")
}

func (ch *ConsoleHandler) UnsetInverse() {
	ch.EmitSequence("27m")
}

func (ch *ConsoleHandler) ReadLine() string {
	line, _ := ch.reader.ReadString('\n')

	if 0 == strings.Compare(runtime.GOOS, "windows") {
		line = strings.Replace(line, "\r\n", "", -1)
	} else {
		line = strings.Replace(line, "\n", "", -1)
	}

	return line
}

func (ch *ConsoleHandler) DisplayTest() {
	ch.ClearScreen()
	ch.CursorToHome()
	
	ch.SetForegroundBlack()
	ch.SetBackgroundWhite()
	fmt.Println("Foreground: Black")
	ch.Reset()
	ch.SetForegroundRed()
	fmt.Println("Foreground: Red")
	ch.SetForegroundGreen()
	fmt.Println("Foreground: Green")
	ch.SetForegroundYellow()
	fmt.Println("Foreground: Yellow")
	ch.SetForegroundBlue()
	fmt.Println("Foreground: Blue")
	ch.SetForegroundMagenta()
	fmt.Println("Foreground: Magenta")
	ch.SetForegroundCyan()
	fmt.Println("Foreground: Cyan")
	ch.SetForegroundWhite()
	fmt.Println("Foreground: White")
	ch.SetForegroundDefault()
	fmt.Println("Foreground: Default")

	ch.Reset()
						
	ch.SetBold()
	fmt.Println("BOLD")
	ch.UnsetBold()

	ch.SetItalic()
	fmt.Println("ITALIC")
	ch.UnsetItalic()												

	ch.SetUnderline()
	fmt.Println("UNDERLINE")
	ch.UnsetUnderline()												

	ch.SetBlink()
	fmt.Println("BLINK")
	ch.UnsetBlink()

	ch.SetInverse()
	fmt.Println("INVERSE")
	ch.UnsetInverse()
}

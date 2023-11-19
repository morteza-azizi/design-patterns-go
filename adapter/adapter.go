package adapter

import "fmt"

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct{}

func (l *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s\n", s)
	println(newMsg)
	return
}

type ModernPrinter interface {
	PrintMessage() string
}

type PrinterAdapter struct {
	HpLegacyPrinter LegacyPrinter
	Message         string
}

func (p *PrinterAdapter) PrintMessage() (newMessage string) {
	if p.HpLegacyPrinter != nil {
		newMessage = fmt.Sprintf("Adapter: %s", p.Message)
		newMessage = p.HpLegacyPrinter.Print(newMessage)
	} else {
		newMessage = p.Message
	}

	return
}

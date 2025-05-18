package structural

import (
	"fmt"
	"strings"
)

type Writer interface {
	Write(s string) string
}

type ConsoleWriter struct{}

func (c *ConsoleWriter) Write(s string) string {
	return fmt.Sprintf("ConsoleWriter: %s", s)
}

type ModerWriter interface {
	WriteText(s string) string
}

type ModernConsultWriter struct{}

func (c *ModernConsultWriter) WriteText(s string) string {
	return fmt.Sprintf("ModernConsultWriter: %s", strings.ToUpper(s))
}

type PrinterAdapter struct {
	printer ModernConsultWriter
}

func (p *PrinterAdapter) Print(s string) string {
	fmt.Println("Adapter translating Print to PrintText")
	return p.printer.WriteText(s)
}

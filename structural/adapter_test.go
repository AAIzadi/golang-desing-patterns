package structural

import (
	"testing"
)

func TestConsoleWriter(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic test",
			input:    "hello",
			expected: "ConsoleWriter: hello",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "ConsoleWriter: ",
		},
	}

	cw := &ConsoleWriter{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cw.Write(tt.input)
			if got != tt.expected {
				t.Errorf("ConsoleWriter.Write() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestModernConsultWriter(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic test",
			input:    "hello",
			expected: "ModernConsultWriter: HELLO",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "ModernConsultWriter: ",
		},
	}

	mcw := &ModernConsultWriter{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mcw.WriteText(tt.input)
			if got != tt.expected {
				t.Errorf("ModernConsultWriter.WriteText() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPrinterAdapter(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic test",
			input:    "hello",
			expected: "ModernConsultWriter: HELLO",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "ModernConsultWriter: ",
		},
	}

	adapter := &PrinterAdapter{
		printer: ModernConsultWriter{},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := adapter.Print(tt.input)
			if got != tt.expected {
				t.Errorf("PrinterAdapter.Print() = %v, want %v", got, tt.expected)
			}
		})
	}
}

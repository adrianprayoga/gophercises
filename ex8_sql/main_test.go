package main

import (
	"testing"
)

func TestParseHtml(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"1234567890", "1234567890"},
		{"123 456 7891", "1234567891"},
		{"(123) 456 7892", "1234567892"},
		{"(123) 456-7893", "1234567893"},
		{"123-456-7894", "1234567894"},
		{"123-456-7890", "1234567890"},
		{"1234567892", "1234567892"},
		{"(123)456-7892", "1234567892"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := CleanPhone(tt.input)
			if output != tt.expected {
				t.Errorf("got %v, want %v", output, tt.expected)
			}
		})
	}
}

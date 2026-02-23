package parse

import (
	"testing"
)

func TestByteToStringParse1(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
	}{
		{"empty", []byte{}},
		{"ascii only", []byte("hello")},
		{"with null block", []byte("foo\x00\x00\x00\x00\x00\x00\x00\x00bar")},
		{"mixed", []byte("a\x01b")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ByteToStringParse1(tt.in)
			if len(tt.in) > 0 && got == "" && !containsNullBlock(tt.in) {
				t.Errorf("ByteToStringParse1 got empty for non-empty input")
			}
		})
	}
}

func TestByteToStringParse2(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		want string
	}{
		{"empty", []byte{}, ""},
		{"ascii only", []byte("hello"), "hello"},
		{"with hex escape", []byte("a\x00b"), "a\\x00b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ByteToStringParse2(tt.in)
			if got != tt.want {
				t.Errorf("ByteToStringParse2(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func containsNullBlock(b []byte) bool {
	nul := []byte("\x00\x00\x00\x00\x00\x00\x00\x00")
	for i := 0; i <= len(b)-len(nul); i++ {
		match := true
		for j := 0; j < len(nul); j++ {
			if b[i+j] != nul[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

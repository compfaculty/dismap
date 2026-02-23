package parse

import (
	"testing"
)

func TestNetJudgeParse(t *testing.T) {
	tests := []struct {
		name  string
		target string
		want  bool
	}{
		{"valid ip", "192.168.1.1", true},
		{"valid ip 2", "10.0.0.1", true},
		{"valid cidr", "192.168.1.0/24", true},
		{"valid cidr 2", "10.0.0.0/8", true},
		{"valid range 1-254", "192.168.1.1-254", true},
		{"valid range 1-10", "192.168.1.1-10", true},
		{"invalid - not ip", "notanip", false},
		{"single ip valid", "1.1.1.1", true},
		{"invalid cidr", "256.256.256.256/24", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NetJudgeParse(tt.target)
			if got != tt.want {
				t.Errorf("NetJudgeParse(%q) = %v, want %v", tt.target, got, tt.want)
			}
		})
	}
}

package parse

import (
	"reflect"
	"testing"
)

func TestPortParse(t *testing.T) {
	tests := []struct {
		name  string
		in    string
		want  []int
		check func([]int) bool
	}{
		{"empty uses default", "", nil, func(got []int) bool { return len(got) > 0 }},
		{"single port", "80", []int{80}, nil},
		{"comma list", "80,443,8080", []int{80, 443, 8080}, nil},
		{"range", "1-5", []int{1, 2, 3, 4, 5}, nil},
		{"range single", "22-22", []int{22}, nil},
		{"range large", "8080-8082", []int{8080, 8081, 8082}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PortParse(tt.in)
			if tt.check != nil {
				if !tt.check(got) {
					t.Errorf("PortParse(%q) = %v (len=%d), check failed", tt.in, got, len(got))
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PortParse(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

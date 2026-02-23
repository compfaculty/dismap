package parse

import (
	"testing"
)

func TestSchemeParse(t *testing.T) {
	tests := []struct {
		name   string
		result map[string]interface{}
		want   string
	}{
		{
			name: "http with path",
			result: map[string]interface{}{
				"path":     "/api",
				"protocol": "http",
				"port":     8080,
				"host":     "example.com",
			},
			want: "http://example.com:8080/api",
		},
		{
			name: "https no path",
			result: map[string]interface{}{
				"path":     "",
				"protocol": "https",
				"port":     443,
				"host":     "example.com",
			},
			want: "https://example.com:443",
		},
		{
			name: "no scheme - host:port",
			result: map[string]interface{}{
				"path":     "",
				"protocol": "",
				"port":     3306,
				"host":     "db.example.com",
			},
			want: "db.example.com:3306",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SchemeParse(tt.result)
			if got != tt.want {
				t.Errorf("SchemeParse() = %q, want %q", got, tt.want)
			}
		})
	}
}

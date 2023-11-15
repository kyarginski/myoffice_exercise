package processes

import (
	"testing"
)

func Test_isValidURL(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    bool
		wantErr bool
	}{
		{
			name:    "Valid URL",
			input:   "https://www.wikipedia.org",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Empty URL",
			input:   "",
			want:    false,
			wantErr: true,
		},
		{
			name:    "Error URL",
			input:   "12345678",
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := isValidURL(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("isValidURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isValidURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

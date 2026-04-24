package internal

import "testing"

func TestExtractField(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{name: "happy path", input: "MemTotal:       32699416 kB", want: "32699416", wantErr: false},
		{name: "empty string", input: "", want: "", wantErr: true},
		{name: "no kB field", input: "HugePages_Free:        0", want: "0", wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractFieldFromLine(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("wantErr %v, got err: %v", tt.wantErr, err)
			}
			if got != tt.want {
				t.Errorf("want %s, got %s", tt.want, got)
			}
		})
	}
}

package internal

import "testing"

func TestExtractField(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{name: "happy path", input: "MemTotal:       32699416 kB", want: 32699416, wantErr: false},
		{name: "empty string", input: "", want: -1, wantErr: true},
		{name: "non numeric val", input: "MemTotal: x kB", want: -1, wantErr: true}, // not sure if this is a realistic use case tbh
		{name: "no kB field", input: "HugePages_Free:        0", want: 0, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractFieldFromLine(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("wantErr %v, got err: %v", tt.wantErr, err)
			}
			if got != tt.want {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}

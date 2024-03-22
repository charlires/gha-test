package cli

import "testing"

func TestDoX(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "x",
			want: "do",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoX(); got != tt.want {
				t.Errorf("DoX() = %v, want %v", got, tt.want)
			}
		})
	}
}

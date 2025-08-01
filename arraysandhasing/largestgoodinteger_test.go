package arraysandhasing

import "testing"

func TestLargestGoodInteger(t *testing.T) {
	tests := []struct {
		num      string
		expected string
	}{
		{"6777133339", "777"},
		{"2300019", "000"},
		{"42352338", ""},
		{"222", "222"},
		{"", ""},
		{"111222333", "333"},
		{"999888777", "999"},
	}

	for _, tt := range tests {
		result := LargestGoodInteger(tt.num)
		if result != tt.expected {
			t.Errorf("LargestGoodInteger(%q) = %q; want %q", tt.num, result, tt.expected)
		}
	}
}

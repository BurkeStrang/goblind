package arraysandhashing

import "testing"

func TestMakeEqual(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  bool
	}{
		{
			name:  "Example 1",
			words: []string{"abc", "aabc", "bc"},
			want:  true,
		},
		{
			name:  "Example 2",
			words: []string{"ab", "a"},
			want:  false,
		},
		{
			name:  "All equal already",
			words: []string{"a", "a", "a"},
			want:  true,
		},
		{
			name:  "Impossible case",
			words: []string{"a", "b", "c"},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeEqual(tt.words); got != tt.want {
				t.Errorf("MakeEqual(%v) = %v, want %v", tt.words, got, tt.want)
			}
		})
	}
}

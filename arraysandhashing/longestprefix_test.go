package arraysandhashing

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
    tests := []struct {
        input    []string
        expected string
    }{
        {[]string{"flower", "flow", "flight"}, "fl"},
        {[]string{"dog", "racecar", "car"}, ""},
        {[]string{"interspecies", "interstellar", "interstate"}, "inters"},
        {[]string{"throne", "throne"}, "throne"},
        {[]string{}, ""},
        {[]string{"a"}, "a"},
        {[]string{"ab", "a"}, "a"},
    }

    for _, tt := range tests {
        result := LongestCommonPrefix(tt.input)
        if result != tt.expected {
            t.Errorf("LongestCommonPrefix(%v) = %q; want %q", tt.input, result, tt.expected)
        }
    }
}

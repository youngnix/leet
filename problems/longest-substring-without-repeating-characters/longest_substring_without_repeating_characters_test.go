package longestsubstringwithoutrepeatingcharacters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{
		{
			s:    "abcabcbb",
			want: 3,
		},
		{
			s:    "bbbbb",
			want: 1,
		},
		{
			s:    "pwwkew",
			want: 3,
		},
		{
			s:    "aab",
			want: 2,
		},
		{
			s:    "dvdf",
			want: 3,
		},
	}

	for _, c := range cases {
		t.Run(c.s, func(t *testing.T) {
			got := lengthOfLongestSubstring(c.s)

			if got != c.want {
				t.Errorf("expected %d, got %d", c.want, got)
			}
		})
	}
}

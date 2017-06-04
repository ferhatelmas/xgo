package anagram

import (
	"fmt"
	"sort"
	"testing"
)

const targetTestVersion = 2

func equal(a []string, b []string) bool {
	if len(b) != len(a) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

func TestDetectAnagrams(t *testing.T) {
	for _, tt := range testCases {
		actual := Detect(tt.subject, tt.candidates)
		if !equal(tt.expected, actual) {
			msg := `FAIL: %s
	Subject %s
	Candidates %q
	Expected %q
	Got %q
				`
			t.Fatalf(msg, tt.description, tt.subject, tt.candidates, tt.expected, actual)
		} else {
			t.Logf("PASS: %s", tt.description)
		}
	}
}

func BenchmarkDetectAnagrams(b *testing.B) {

	b.StopTimer()

	for _, tt := range testCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			Detect(tt.subject, tt.candidates)
		}

		b.StopTimer()
	}

}

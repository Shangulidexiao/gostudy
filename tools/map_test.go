package tools

import (
	"math/rand"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	t.Error("your map is default")
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < 10; i++ {

	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"abcba", true},
		{"abcbc", false},
		{"ffdfdafa", false},
		{"fafafa", false},
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q)=%v", test.input, got)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	tests := []struct {
		input string
		want  bool
	}{
		{RandomPalindrome(rng), true},
		// {RandomNotPalindrome(rng), false},
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			b.Errorf("IsPalindrome(%q)=%v", test.input, got)
		}
	}
}

func RandomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	count := (n + 1) / 2
	runes := make([]rune, n)
	for i := 0; i < count; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func RandomNotPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		r1 := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r1
	}
	return string(runes)
}

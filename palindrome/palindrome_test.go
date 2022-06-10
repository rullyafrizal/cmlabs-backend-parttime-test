package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome("racecar"))
	assert.False(t, IsPalindrome("rudy"))
	assert.True(t, IsPalindrome("SAIPPUAKIVIKAUPPIAS"))
	assert.True(t, IsPalindrome("Anna"))
	assert.True(t, IsPalindrome("Civic"))
	assert.True(t, IsPalindrome("My gym"))
	assert.True(t, IsPalindrome("No lemon, nomelon"))
}

func BenchmarkIsPalindrome(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsPalindrome("No lemon, nomelon")
	}
}

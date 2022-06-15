package pow

import (
	"bytes"
	"encoding/hex"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {

	const difficulty = 0x1000000

	token := NewToken(difficulty)

	assert.True(t, strings.HasPrefix(hex.EncodeToString(token), "000000"))
}

func TestSolve(t *testing.T) {

	const difficulty = 0x10000

	rand.Seed(1) // generate not random token
	token := NewToken(difficulty)

	nonce := Solve(token)

	assert.True(t, Verify(token, nonce))
	assert.True(t, bytes.Compare(hash(token, nonce), token) < 0)
	assert.True(t, hex.EncodeToString(nonce) > "0000000000001000")
}

func BenchmarkSolve(b *testing.B) {
	const difficulty = 0x100
	token := NewToken(difficulty)

	for i := 0; i < b.N; i++ {
		Solve(token)
	}
}

func BenchmarkVerify(b *testing.B) {
	b.StopTimer()
	const difficulty = 0x100
	token := NewToken(difficulty)
	nonce := Solve(token)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Verify(token, nonce)
	}
}

package pow

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewToken generates random token for PoW with givens difficulty
func NewToken(difficulty uint64) []byte {
	buf := make([]byte, 16)
	binary.BigEndian.PutUint64(buf[:8], math.MaxUint64/difficulty)
	rand.Read(buf[8:])
	return buf
}

// Solve solves PoW-challenge for givens token
func Solve(token []byte) (nonce []byte) {
	nonce = make([]byte, 8)
	for i := uint64(0); ; i++ {
		binary.BigEndian.PutUint64(nonce, i)
		if Verify(token, nonce) {
			return
		}
	}
}

// Verify verifies nonce for givens token
func Verify(token, nonce []byte) bool {
	hash := hash(token, nonce)
	return bytes.Compare(hash, token) < 0
}

// hash-function of PoW
func hash(token, nonce []byte) []byte {
	h := sha256.New()
	h.Write(token)
	h.Write(nonce)
	return h.Sum(nil)
}

package util

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func NewUlid() string {
	t := time.Now()
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	// entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), order)
	entropy := ulid.Monotonic(rand.New(rand.NewSource(seed.Int64())), 0)
	ulid := ulid.MustNew(ulid.Timestamp(t), entropy).String()
	return ulid
}

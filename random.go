package random

import (
	crand "crypto/rand"
	"math/big"
	mrand "math/rand"
	"strings"
	"time"

	"github.com/hashicorp/go-uuid"
)

type RNG interface {
	Intn(max int) int
}

type defaultRNG struct {
	r *mrand.Rand
}

func NewDefaultRNG() RNG {
	return &defaultRNG{
		r: mrand.New(mrand.NewSource(time.Now().UnixNano())),
	}
}

func (s *defaultRNG) Intn(max int) int {
	return s.r.Intn(max)
}

type secureRNG struct{}

func NewSecureRNG() RNG {
	return &secureRNG{}
}

func (s *secureRNG) Intn(max int) int {
	val, _ := crand.Int(crand.Reader, big.NewInt(int64(max)))
	return int(val.Int64())
}

type Alphabet struct {
	source []rune
}

func (s Alphabet) Gen(n int) string {
	return s.RGen(NewDefaultRNG(), n)
}

func (s Alphabet) RGen(rng RNG, n int) string {
	max := len(s.source)
	result := make([]rune, n)
	for i := 0; i < n; i++ {
		pos := rng.Intn(max)
		result[i] = s.source[pos]
	}
	return string(result)
}

const (
	lower       = "abcdefghijklmnopqrstuvwxyz"
	upper       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits      = "0123456789"
	punctuation = ".:,:"
	accents     = "´`^"
	dashes      = "_-"
	slashes     = "/\\"
	brackets    = "()[]{}<>"
	whitespace  = "\t \v"
	printable   = "!\"#$%&'()*+,-./" + digits + ":;<=>?@" + upper + "[\\]^_`" + lower + "{|}~"
	hex         = "0123456789ABCDEF"
	urlSafe     = lower + upper + digits + "-._~"
)

func fromString(s ...string) Alphabet {
	return Alphabet{
		source: []rune(strings.Join(s, "")),
	}
}

func Hex() Alphabet {
	return fromString(hex)
}

func String() Alphabet {
	return fromString(digits, lower, upper)
}

func Printable() Alphabet {
	return fromString(printable)
}

func Digits() Alphabet {
	return fromString(digits)
}

func URLSafe() Alphabet {
	return fromString(urlSafe)
}

func GenUUID() string {
	result, _ := uuid.GenerateUUID()
	return result
}

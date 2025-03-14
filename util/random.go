package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var seededRand *rand.Rand

func init() {
    seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandInt returns a random integer between min and max
func RandInt(min, max int64) int64 {
    return min + rand.Int63n(max-min+1) 
}

// RandString returns a random string of length n
func RandString(length int) string {
    var sb strings.Builder
    k := len(alphabet)

    for i := 0; i < length; i++ {
        c := alphabet[seededRand.Intn(k)]
        sb.WriteByte(c)
    }

    return sb.String()
}

// RandOwner returns a random owner name
func RandOwner() string {
    return RandString(6)
}

// RandMoney returns a random amount of money
func RandMoney() int64 {
    return RandInt(0, 1000)
}

// RandCurrency returns a random currency
func RandCurrency() string {
    currencies := []string{"USD", "EUR", "CAD"}
    n := len(currencies)
    return currencies[rand.Intn(n)]
}

// RandomEmail generates a random email
func RandEmail() string {
	return fmt.Sprintf("%s@email.com", RandString(6))
}

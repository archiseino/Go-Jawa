package token

import (
	"time"
)

type Maker interface {
	// Consist of two methods signature for specific token and duration
	CreateToken(username string, role string, duration time.Duration) (string, *Payload, error)

	VerifyToken(token string) (*Payload, error)
}

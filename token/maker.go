package token

import (
	"time"
)

type Maker interface {
	// Consist of two methods signature for specific token and duration
	CreateToken(username string, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}

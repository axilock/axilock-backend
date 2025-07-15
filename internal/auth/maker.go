package auth

import "time"

type Maker interface {
	CreateToken(userID string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}

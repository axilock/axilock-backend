package auth

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrExpiredToken = errors.New("token is expired")

type Payload struct {
	ID        uuid.UUID `json:"id,omitempty"`
	UserID    string    `json:"userid,omitempty"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(userID string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        tokenID,
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	// if time.Now().After(payload.ExpiredAt) {
	// 	return ErrExpiredToken
	// } // TODO: remove validity check/ put after refresh token
	return nil
}

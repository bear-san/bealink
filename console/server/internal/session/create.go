package session

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/bear-san/bealink/console/server/pkg/random_string"
	"time"
)

func Create(ctx context.Context, s *Session) (*string, error) {
	token, err := random_string.Create(32)
	if err != nil {
		return nil, err
	}
	s.ID = &token

	hashedToken := sha256.Sum256([]byte(token))
	sessionID := base64.StdEncoding.EncodeToString(hashedToken[:])

	payload, err := json.Marshal(s)
	if err != nil {
		return nil, err

	}

	return &token, client.Set(ctx, fmt.Sprintf("bealink-session-%s", sessionID), (string)(payload), time.Until(s.ExpiresAt)).Err()
}

type Session struct {
	ID        *string   `json:"id"`
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	ExpiresAt time.Time `json:"expires_at"`
}

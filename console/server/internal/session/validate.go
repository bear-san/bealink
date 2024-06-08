package session

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func ExtractToken(req *gin.Context) *string {
	token, err := req.Cookie("token")
	if err == nil && token != "" {
		return &token
	}

	header := req.GetHeader("Authorization")
	if header == "" {
		return nil
	}

	hValues := strings.Split(header, " ")
	if len(hValues) != 2 || hValues[0] != "Bearer" {
		return nil
	}

	return &hValues[1]
}

func Validate(ctx context.Context, token string) (*Session, error) {
	hashedToken := sha256.Sum256([]byte(token))
	sessionID := base64.StdEncoding.EncodeToString(hashedToken[:])

	payload, err := client.Get(ctx, fmt.Sprintf(sessionKeyFormat, sessionID)).Result()
	if err != nil {
		return nil, err
	}

	var s Session
	if err := json.Unmarshal([]byte(payload), &s); err != nil {
		return nil, err
	}

	if s.ExpiresAt.Before(time.Now()) {
		return nil, fmt.Errorf("session expired")
	}

	return &s, nil
}

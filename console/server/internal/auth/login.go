package auth

import (
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"os"
)

func Login(req *gin.Context) {
	provider, err := oidc.NewProvider(req.Request.Context(), os.Getenv("OIDC_PROVIDER"))
	if err != nil {
		req.JSON(500, gin.H{"error": err.Error()})
		return
	}

	scheme := "http"
	if req.Request.TLS != nil {
		scheme = "https"
	}

	config := oauth2.Config{
		ClientID:     os.Getenv("OIDC_CLIENT_ID"),
		ClientSecret: os.Getenv("OIDC_CLIENT_SECRET"),
		Endpoint:     provider.Endpoint(),
		RedirectURL:  fmt.Sprintf("%s://%s/api/auth/callback", scheme, req.Request.Host),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	req.Redirect(302, config.AuthCodeURL("", oauth2.AccessTypeOffline))
}

package auth

import (
	"fmt"
	"github.com/bear-san/bealink/console/server/internal/session"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"math"
	"os"
	"time"
)

func Callback(req *gin.Context) {
	code := req.Query("code")
	if code == "" {
		req.JSON(400, gin.H{"error": "code is required"})
		return
	}

	client, err := oidc.NewProvider(req.Request.Context(), os.Getenv("OIDC_PROVIDER"))
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
		Endpoint:     client.Endpoint(),
		RedirectURL:  fmt.Sprintf("%s://%s/api/auth/callback", scheme, req.Request.Host),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	token, err := config.Exchange(req.Request.Context(), code)
	if err != nil {
		req.JSON(500, gin.H{"error": err.Error()})
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		req.JSON(500, gin.H{"error": "missing id_token"})
		return
	}

	user, err := client.Verifier(&oidc.Config{ClientID: os.Getenv("OIDC_CLIENT_ID")}).Verify(req.Request.Context(), rawIDToken)
	if err != nil {
		req.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var claims struct {
		Email      string `json:"email"`
		UserName   string `json:"preferred_username"`
		FamilyName string `json:"family_name"`
		GivenName  string `json:"given_name"`
	}

	if err := user.Claims(&claims); err != nil {
		req.JSON(500, gin.H{"error": err.Error()})
		return
	}

	expireAt := time.Now().Add(time.Hour * 3)

	t, err := session.Create(req.Request.Context(), &session.Session{
		UserID:    claims.UserName,
		Email:     claims.Email,
		FullName:  fmt.Sprintf("%s %s", claims.FamilyName, claims.GivenName),
		ExpiresAt: expireAt,
	})

	req.SetCookie("token", *t, (int)(math.Round(time.Until(expireAt).Seconds())), "/", "", false, true)
	req.Redirect(302, "/")
}

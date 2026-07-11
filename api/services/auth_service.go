package services

import (
	"context"
	"fmt"
	"log"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"golang.org/x/crypto/bcrypt"

	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbquery "github.com/SONEsee/go-echo/pkg/db-pkg/db-query"
	"github.com/SONEsee/go-echo/pkg/jwt-pkg/auth"
)

func LoginServices(ctx context.Context, req requestbody.UserLoginRequest) (*requestbody.UserLoginResponse, error) {
	user, err := dbquery.GetUserByUsername(ctx, dbpkg.DB, req.UserName)
	if err != nil {
		log.Printf("❌ User not found: %s", req.UserName)
		return nil, fmt.Errorf("invalid username or password")
	}

	if !user.IsActive {
		log.Printf("❌ User blocked: %s", req.UserName)
		return nil, fmt.Errorf("account has been blocked")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		log.Printf("❌ Password mismatch for user: %s", req.UserName)
		return nil, fmt.Errorf("invalid username or password")
	}

	token, err := auth.GenerateToken(int64(user.ID), user.Username, user.RoleID)
	if err != nil {
		log.Printf("❌ Failed to generate token: %v", err)
		return nil, fmt.Errorf("failed to generate token")
	}

	log.Printf("✅ Login successful for user: %s", user.Username)

	return &requestbody.UserLoginResponse{
		ID:       int64(user.ID),
		FullName: user.FullName,
		UserName: user.Username,
		Email:    user.Email,
		RoleID:   user.RoleID,
		Token:    token,
	}, nil
}

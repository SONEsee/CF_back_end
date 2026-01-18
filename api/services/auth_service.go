package services

import (
	"context"
	"fmt"
	"log"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"golang.org/x/crypto/bcrypt"

	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbquery "github.com/SONEsee/go-echo/pkg/db-pkg/db-query"
	"github.com/SONEsee/go-echo/pkg/jwt-pkg/auth" // ✅ Import jwt
)

func LoginServices(ctx context.Context, req requestbody.UserLoginRequest) (*requestbody.UserLoginResponse, error) {
	// ດຶງຂໍ້ມູນ user
	user, err := dbquery.GetUserByUsername(ctx, dbpkg.DB, req.UserName)
	if err != nil {
		log.Printf("❌ User not found: %s", req.UserName)
		return nil, fmt.Errorf("invalid username or password")
	}

	// ເຊັກ blacklist
	if user.BackList {
		log.Printf("❌ User blacklisted: %s", req.UserName)
		return nil, fmt.Errorf("account has been blocked")
	}

	// ເຊັກ password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Printf("❌ Password mismatch for user: %s", req.UserName)
		return nil, fmt.Errorf("invalid username or password")
	}

	// ✅ ສ້າງ JWT token
	token, err := auth.GenerateToken(int64(user.ID), user.UserName, user.RoleID)
	if err != nil {
		log.Printf("❌ Failed to generate token: %v", err)
		return nil, fmt.Errorf("failed to generate token")
	}

	log.Printf("✅ Login successful for user: %s", user.UserName)

	// Return response
	response := &requestbody.UserLoginResponse{
		ID:         int64(user.ID),
		Name:       user.Name,
		FullName:   user.FullName,
		UserName:   user.UserName,
		ProfileImg: user.ProfileImg,
		RoleID:     user.RoleID,
		Token:      token, // ✅ ສົ່ງ token ກັບໄປ
	}

	return response, nil
}

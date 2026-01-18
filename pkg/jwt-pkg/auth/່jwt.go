package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ✅ ປ່ຽນ secret key ນີ້ໃນ production!
var jwtSecret = []byte("your-super-secret-key-change-this")

type JWTClaims struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	RoleID   int    `json:"role_id"`
	jwt.RegisteredClaims
}

// GenerateToken ສ້າງ JWT token
func GenerateToken(userID int64, username string, roleID int) (string, error) {
	claims := JWTClaims{
		UserID:   userID,
		UserName: username,
		RoleID:   roleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24 ຊົ່ວໂມງ
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

// ValidateToken ກວດສອບ token
func ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// ເຊັກວ່າໃຊ້ signing method ຖືກຕ້ອງ
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		// ເຊັກວ່າໝົດອາຍຸບໍ່
		if claims.ExpiresAt.Time.Before(time.Now()) {
			return nil, fmt.Errorf("token expired")
		}
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// GetSecretKey ດຶງ secret key (ສຳລັບການຕັ້ງຄ່າ)
func GetSecretKey() []byte {
	return jwtSecret
}

// SetSecretKey ຕັ້ງຄ່າ secret key (ສຳລັບການຕັ້ງຄ່າຈາກ env)
func SetSecretKey(key string) {
	jwtSecret = []byte(key)
}

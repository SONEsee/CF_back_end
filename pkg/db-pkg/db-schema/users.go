package dbschema

import "time"

type UserDBSchema struct {
	ID           int        `db:"id" json:"id"`
	ShopID       *int       `db:"shop_id" json:"shop_id"`
	RoleID       int        `db:"role_id" json:"role_id"`
	Username     string     `db:"username" json:"username"`
	PasswordHash string     `db:"password_hash" json:"-"`
	FullName     string     `db:"full_name" json:"full_name"`
	Email        string     `db:"email" json:"email"`
	Phone        string     `db:"phone" json:"phone"`
	ProfileImage *string    `db:"profile_image" json:"profile_image"`
	IsActive     bool       `db:"is_active" json:"is_active"`
	LastLoginAt  *time.Time `db:"last_login_at" json:"last_login_at"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updated_at"`
}

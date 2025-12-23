package requestbody

import "time"

type Role struct {
	ID        int       `json:"id"`
	RoleName  string    `json:"role_name" validate:"required,min=2,max=150"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
type RolePatchRequest struct {
	RoleName *string `json:"role_name,omitempty" validate:"omitempty,min=2,max=150"`
	Detail   *string `json:"detail,omitempty"`
}

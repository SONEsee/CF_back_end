package dbschema

type GetUserDataDBSchema struct {
	ID         int    `json:"id"`
	Name       string `db:"name" json:"name"`
	FullName   string `db:"full_name" json:"full_name"`
	UserName   string `db:"user_name" json:"user_name"`
	Password   string `db:"password" json:"password"`
	ProfileImg string `db:"profile_image" json:"profile_image"`
	BackList   bool   `db:"back_list" json:"black_list"`
	RoleID     int    `db:"role_id" json:"role_id"`
}

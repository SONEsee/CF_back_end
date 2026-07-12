package dbschema

import "time"

type LiveSessionDBSchema struct {
	ID              int        `db:"id" json:"id"`
	SocialAccountID int        `db:"social_account_id" json:"social_account_id"`
	FbVideoID       *string    `db:"fb_video_id" json:"fb_video_id"`
	SessionTitle    *string    `db:"session_title" json:"session_title"`
	Status          string     `db:"status" json:"status"`
	StartedAt       time.Time  `db:"started_at" json:"started_at"`
	EndedAt         *time.Time `db:"ended_at" json:"ended_at"`
}

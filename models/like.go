package models

import "github.com/google/uuid"

type Like struct {
	ID      uint      `json:"id"`
	PostID  uint      `json:"post_id"`
	UserID  uuid.UUID `json:"user_id"`
	LikedAt string    `json:"liked_at"`
}

package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID           uint      `json:"id" db:"id"`
	PostID       uint      `json:"post_id" db:"post_id"`
	UserID       uuid.UUID `json:"user_id" db:"user_id"`
	CommentText  string    `json:"comment_text" db:"comment_text"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type CommentWithUser struct {
	Comment
	Username  string `json:"username" db:"username"`
	ProfilePic string `json:"profile_pic" db:"profile_pic"`
}
package models

import "github.com/google/uuid"

type Role string

const (
    RoleUser  Role = "user"
    RoleAdmin Role = "admin"
)

type User struct {
    ID         uuid.UUID  `json:"id"          gorm:"type:uuid;primaryKey"`
    Username   string     `json:"username"    gorm:"type:varchar(50);uniqueIndex;not null"`
    Name       string     `json:"name"        gorm:"type:varchar(100);not null"`
    ProfilePic *string    `json:"profile_pic" gorm:"type:text"`
    Role       Role       `json:"role"        gorm:"type:varchar(10);not null"`
    Bio        *string    `json:"bio"         gorm:"type:text"`
    Email      *string    `json:"email"       gorm:"type:text"`
}
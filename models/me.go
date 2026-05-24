package models

type Me struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	ProfilePic string `json:"profile_pic"`
	Role       string `json:"role"`
}

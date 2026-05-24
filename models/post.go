package models

type Post struct {
	ID          uint   `json:"id" db:"id"`
	Image       string `json:"image" db:"image"`
	CategoryId uint   `json:"category_id" db:"category_id"`
	CategoryName string `json:"category_name" db:"category_name"`
	Title       string `json:"title" db:"title"`
	Description *string `json:"description" db:"description"`
	Date        string `json:"date" db:"date"`
	Content     string `json:"content" db:"content"`
	StatusId   uint   `json:"status_id" db:"status_id"`
	LikeCount   int    `json:"likes_count" db:"likes_count"`
}

type PostResponse struct {
	TotalPosts  int    `json:"totalPosts"`
	TotalPages  int    `json:"totalPages"`
	CurrentPage int    `json:"currentPage"`
	Limit       int    `json:"limit"`
	Posts       []Post `json:"posts"`
	NextPage    *int   `json:"nextPage"`
}

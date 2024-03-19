package entity

import "time"

type (
	Creator struct {
		UserId      int    `json:"userId"`
		Name        string `json:"name"`
		ImageUrl    string `json:"imageUrl"`
		FriendCount int    `json:"friendCount"`
		CreatedAt   string `json:"createdAt"`
	}

	CommentPerPost struct {
		Comment string  `json:"comment"`
		Creator Creator `json:"creator"`
	}

	Post struct {
		Id         int              `json:"id"`
		PostInHtml string           `json:"postInHtml"`
		Tags       []string         `json:"tags"`
		UserID     int              `json:"userId"`
		CreatedAt  time.Time        `json:"createdAt"`
		Comments   []CommentPerPost `json:"comments"`
	}
)

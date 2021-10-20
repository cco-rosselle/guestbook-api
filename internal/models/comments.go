package models

type Comment struct {
	Description string `json:"description" bson:"description"`
	CommentID string `json:"commentId" bson:"commentId"`
}

type Comments struct {
	Data []*Comment `json:"data"`
	Total int `json:"total"`
}
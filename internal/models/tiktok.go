package models

type TikTokVideo struct{
	Username     string `json:"username"`      
	Caption      string `json:"caption"`      
	VideoURL     string `json:"video_url"`     
	LikeCount    int    `json:"like_count"`   
	CommentCount int    `json:"comment_count"` 
}
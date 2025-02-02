package models

type TikTokVideo struct{
	Username     string `json:"username"`      // TikTok username
	Caption      string `json:"caption"`       // Video caption
	VideoURL     string `json:"video_url"`     // Direct video URL
	LikeCount    int    `json:"like_count"`    // Number of likes
	CommentCount int    `json:"comment_count"` // Number of comments
}
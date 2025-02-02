package models

// FacebookPost represents the metadata for a Facebook post
type FacebookPost struct {
	Title       string `json:"title"`        // Title of the post
	Description string `json:"description"`  // Post description
	ImageURL    string `json:"image_url"`    // URL of the featured image
}

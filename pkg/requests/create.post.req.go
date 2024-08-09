package requests

type CreatePostReq struct {
	TextContent string `json:"textContent"`
	ImageContent string `json:"imageContent"`
}
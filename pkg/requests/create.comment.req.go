package requests

type CreateCommentReq struct {
	PostID  int    `json:"postID"`
	Content string `json:"content"`
}

package requests

type UpdateCommentReq struct {
	CommentID int    `json:"commentID"`
	Content   string `json:"content"`
}

package schema

type ApiResponsePosts struct {
	Code    int    `example:"200"`
	Message string `example:"success"`
	Data    []PostResponse
}

type ApiResponsePost struct {
	Code    int    `example:"200"`
	Message string `example:"success"`
	Data    PostResponse
}

type PostRequest struct {
	Title    string `json:"title" binding:"required,min=20"`
	Content  string `json:"content" binding:"required,min=200"`
	Category string `json:"category" binding:"required,min=3"`
	Status   string `json:"status" binding:"required,oneof='publish' 'draft' 'trash' "`
}

type PostResponse struct {
	Title    string
	Content  string
	Category string
	Status   string
}

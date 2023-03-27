package dto

type TodoRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type TodoResponse struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  string `json:"userId"`
}

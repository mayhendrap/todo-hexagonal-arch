package domain

const (
	TodoCollection = "todos"
)

type Todo struct {
	ID      string `json:"id" bson:"_id,omitempty"`
	Title   string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
	UserID  string `json:"user_id" bson:"user_id"`
}

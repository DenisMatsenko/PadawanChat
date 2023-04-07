package domain

type Message struct {
	Id       int32  `json:"id"`
	Content  string `json:"content"`
	AuthorId int32  `json:"authorId"`
}

package domain

type Message struct {
	Id 			int32	`json:"id"`
	Content 	string	`json:"content"`
	AuthorId 	string	`json:"authorId"`
}


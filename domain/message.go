package domain

type Message struct {
	Id 			int32	`json:"id"`
	Content 	string	`json:"content"`
	Author 		string	`json:"author"`
}


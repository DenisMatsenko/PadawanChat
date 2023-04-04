package domain

type Message struct {
	Id 			int		`json:"id"`
	Content 	string	`json:"content"`
	Author 		string	`json:"author"`
}


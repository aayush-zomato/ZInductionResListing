package models

type Post struct {
	ID		int		`json:"id"`
	Name	string	`json:"name"`
	Cusine	string	`json:"cusine"`
	Distance	string	`json:"distance"`
	Address	string	`json:"address"`
	CFT		string	`json:"cft"`
	Rating	string	`json:"rating"`
	Timing	string	`json:"timing"`
}

// {
// 	"name":"Shop18",
// 	"cusine":"pasta",
// 	"distance":"3km",
// 	"address":"G3S near DTU",
// 	"cft":"Rs. 350",
// 	"rating":"4.0",
// 	"timing":"4 P.M."
// }
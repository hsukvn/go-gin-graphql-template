package model

type User struct {
	ID        int32   `json:"id"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Roles     []*Role `json:"roles"`
	Deposit   int64   `json:"deposit"`
}

// FIXME: Mockup data
var UserData = []User{
	User{
		ID:        1,
		Firstname: "Kevin",
		Lastname:  "Hsu",
		Roles: []*Role{
			&Role{
				ID:   1,
				Name: "Research and Design",
			},
		},
		Deposit: 2147483648,
	},
}

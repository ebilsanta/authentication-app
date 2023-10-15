package internal

type User struct {
    Company    	string
    Id   		string
	Email		string
	FirstName	string
	LastName    string
    Status   	string
	Birthdate	string
}

type Credential struct {
	Company		string	`json:"company"`
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}
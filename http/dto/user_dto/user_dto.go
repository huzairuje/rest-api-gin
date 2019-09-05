package user_dto

type CreateUserDTO struct {
	Name 		string		`json:"name"`
	Email 		string 		`json:"email"`
	Phone 		string 		`json:"phone"`
	Password 	string 		`json:"password"`
	ClientType 	string 		`json:"client_type"`
}

type UpdateUserDTO struct {
	Name 		string		`json:"name"`
	Email 		string 		`json:"email"`
	Phone 		string 		`json:"phone"`
	Password 	string 		`json:"password"`
	ClientType 	string 		`json:"client_type"`
}
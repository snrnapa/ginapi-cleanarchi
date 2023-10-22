package dto

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func NewUserByJson(id string, firstname string, lastname string) *User {
	return &User{
		Id:        id,
		FirstName: firstname,
		LastName:  lastname,
	}

}

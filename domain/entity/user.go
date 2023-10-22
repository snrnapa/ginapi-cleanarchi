package entity

type User struct {
	Id        string `xorm:"id"`
	FirstName string
	LastName  string
}

package entity

type User struct {
	ID       int
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

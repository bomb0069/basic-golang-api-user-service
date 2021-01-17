package user

type Users []User

type User struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Title     string `json:"title"`
}

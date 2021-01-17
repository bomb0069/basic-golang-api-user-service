package user

type Users []User

type User struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Age  int    `json:"age" bson:"age"`
}

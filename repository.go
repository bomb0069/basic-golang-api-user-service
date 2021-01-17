package user

type UserRepository interface {
	GetAll() (Users, error)
}

func NewUserRepository() UserRepository {
	return &MongoUserRepository{}
}

type MongoUserRepository struct {
}

func (repo *MongoUserRepository) GetAll() (Users, error) {

	u := Users{
		User{
			Firstname: "f1",
			Lastname:  "l1",
			Title:     "Mr.",
		},
		User{
			Firstname: "f2",
			Lastname:  "l2",
			Title:     "Miss.",
		},
	}
	return u, nil
}

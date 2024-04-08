package app

// User entity
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type mockDB interface {
	getUsers() []User
}

func (u *User) getUsers() []User {
	return []User{
		{
			ID:   1,
			Name: "Jhon",
			Role: "developer",
		},
		{
			ID:   2,
			Name: "Mariia",
			Role: "developer",
		},
		{
			ID:   3,
			Name: "Silver",
			Role: "admin",
		},
		{
			ID:   4,
			Name: "Jhon",
			Role: "admin",
		},
		{
			ID:   5,
			Name: "Stephan",
			Role: "manager",
		},
	}
}

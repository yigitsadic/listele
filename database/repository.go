package database

type Person struct {
	FullName string `json:"full_name"`
}

type Repository interface {
	FindAll() ([]Person, error)
}

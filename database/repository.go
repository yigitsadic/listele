package database

// Person represents database row for people table.
type Person struct {
	FullName string `json:"full_name"`
}

// Repository is an interface for communicating between handler and database
type Repository interface {
	FindAll() ([]Person, error)
}

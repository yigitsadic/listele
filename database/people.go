package database

import "database/sql"

type PeopleRepository struct {
	Database *sql.DB
}

func (p *PeopleRepository) FindAll() ([]Person, error) {
	rows, err := p.Database.Query("SELECT fullname FROM people")
	if err != nil {
		return nil, err
	}

	var people []Person

	for rows.Next() {
		var person Person

		if err = rows.Scan(&person.FullName); err == nil {
			people = append(people, person)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return people, err
}

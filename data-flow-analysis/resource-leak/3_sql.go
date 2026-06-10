package resource_leak

import "database/sql"

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserNamesByCountry(country string) ([]string, error) {
	rows, err := s.db.Query(`SELECT name FROM users WHERE country = $1`, country) // resource leak!
	if err != nil {
		return nil, err
	}

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			// the rows are not closed here
			return nil, err
		}
		names = append(names, name)
	}

	rows.Close()

	return names, rows.Err()
}

package repository

import "database/sql"

type DB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

func NewMockDB() MockDB {
	return MockDB{
		MockQuery: func(query string, args ...interface{}) (*sql.Rows, error) {
			return nil, nil
		},
	}
}

type MockDB struct {
	MockQuery func(query string, args ...interface{}) (*sql.Rows, error)
}

func (m MockDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return m.MockQuery(query, args...)
}

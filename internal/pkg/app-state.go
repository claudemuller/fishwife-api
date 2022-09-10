package pkg

import "database/sql"

type AppState struct {
	DB *sql.DB
}

func NewAppState() AppState {
	return AppState{
		DB: NewDB(),
	}
}

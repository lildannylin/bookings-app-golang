package dbrepo

import (
	"database/sql"
	"github.com/lildannylin/bookings-app-golang/internal/config"
	"github.com/lildannylin/bookings-app-golang/internal/repository"
)

type postgreDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgreDBRepo{
		App: a,
		DB:  conn,
	}
}

func NewTestingRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &postgreDBRepo{
		App: a,
	}
}

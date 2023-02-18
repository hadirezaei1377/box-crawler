package server

import (
	"box-crawler/database"

	"github.com/labstack/echo/v4"
)

type server struct {
	db   database.Database
	echo *echo.Echo
}

func New(db database.Database) *server {
	return &server{
		db:   db,
		echo: echo.New(),
	}
}

func (srv *server) RegisterRoutes() {
	srv.echo.GET("", SampleHandler(srv.db))
}

func SampleHandler(db database.Database) func(c echo.Context) error {
	return func(c echo.Context) error {
		return nil
	}
}

func (srv *server) Start() error {
	return srv.echo.Start(":8080")
}

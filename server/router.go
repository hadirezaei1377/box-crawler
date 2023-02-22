package server

import (
	"box-crawler/database"
	"net/http"

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
	srv.echo.GET("/upcoming_fights", srv.GetUpcomingFights())
	srv.echo.GET("/upcoming_fights", srv.GetUpcomingFights())
}

func (srv *server) GetUpcomingFights() func(c echo.Context) error {
	return func(c echo.Context) error {
		upcoming_fights, err := srv.db.GetUpcomingFights()
		if err != nil {
			c.Error(err)
		}
		return c.JSON(http.StatusOK, upcoming_fights)
	}
}

func (srv *server) Start() error {
	return srv.echo.Start(":8080")
}

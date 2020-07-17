package web

import (
	"github.com/hsmtkk/clean_arch_study_2/pkg/infra/handler"
	"github.com/hsmtkk/clean_arch_study_2/pkg/interface/controller"
	"github.com/hsmtkk/clean_arch_study_2/pkg/usecase"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	hdl := handler.New(controller.New(usecase.New()))
	e.POST("/users", hdl.CreateUser)
	e.GET("/users", hdl.GetUsers)
	e.Logger.Fatal(e.Start(":8000"))
}

package router

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "github.com/nogizaka-api/userInterface/handler"
  "github.com/labstack/gommon/log"
)

func New() *echo.Echo {
  e := echo.New()

  e.Logger.SetLevel(log.DEBUG)
  e.Pre(middleware.RemoveTrailingSlash())
  e.Use(middleware.Logger())
  e.Use(middleware.CORS())

  member := handler.NewMember()
  g := e.Group("/nogizaka")
  g.GET("/member", member.SelectMember)
  g.GET("/member/all", member.SelectAllMember)

  e.Start(":8000")
  return e
}
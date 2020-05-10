package handler

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/nogizaka-api/userInterface/form"
	"github.com/nogizaka-api/userInterface/response"
	"github.com/nogizaka-api/usecase"
)

type (
	IMember interface {
		SelectMember(c echo.Context) error
		SelectAllMember(c echo.Context) error
	}
	Member struct {
		MemberUsecase usecase.IMember
	}
)

func NewMember() IMember {
	return &Member{
	  MemberUsecase: usecase.NewMember(),
  }
}

func (handler *Member) SelectMember(c echo.Context) error {
	params, err := form.GetParams(c)
	if err != nil {
		return err
	}
	member, err := handler.MemberUsecase.CallSelectMemberByName(params)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response.NewMember(member))
}

func (handler *Member) SelectAllMember(c echo.Context) error {
	params, err := form.GetParams(c)
	if err != nil {
		return err
	}
	members, err := handler.MemberUsecase.CallSelectAllMember(params)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response.NewMembers(members))
}
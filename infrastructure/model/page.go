package model

import "github.com/labstack/echo"

type Page struct {
	Page    int64 `json:"page" validate:"required"`
	PerPage int64 `json:"per" validate:"required"`
}

func InitPage(page int64, perPage int64) *Page {
	return &Page{Page: page, PerPage: perPage}
}

func NewPage(c echo.Context) (*Page, error) {
	form := InitPage(1, 30)
	if err := c.Bind(form); err != nil {
		return nil, err
	}

	if err := c.Validate(form); err != nil {
		return nil, err
	}

	return form, nil

}

func NewPageWithPerPage(c echo.Context, perPage int64) (*Page, error) {
	form := InitPage(1, perPage)
	if err := c.Bind(form); err != nil {
		return nil, err
	}

	if err := c.Validate(form); err != nil {
		return nil, err
	}

	return form, nil
}

func NewPageWithPageAndPerPage(c echo.Context, page int64, perPage int64) (*Page, error) {
	form := InitPage(page, perPage)
	if err := c.Bind(form); err != nil {
		return nil, err
	}

	if err := c.Validate(form); err != nil {
		return nil, err
	}

	return form, nil
}

package controllers

import (
	"github.com/revel/revel"
)

type Calendar struct {
	*revel.Controller
}

func (c Calendar) Calendar() revel.Result {
	return c.Render()
}

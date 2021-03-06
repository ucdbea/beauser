package controllers

import (
	"github.com/revel/revel"
)

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {
	return c.Render()
}

func (c Application) Homepage() revel.Result{
	return c.Render()
}

func (c Application) Signin() revel.Result{
	return c.Render()
}

func (c Application) Signup() revel.Result{
	return c.Render()
}

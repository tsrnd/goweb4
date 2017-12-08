package controllers

import (

)

//  UsersController operations for Users
type UserController struct {
	ExtendController
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Show", c.Show)
}

// @router /users/:id [get]
func (c *UserController) Show() {
	InitFrontEndTemplate(&c.Controller, "frontend/about.tpl")
}

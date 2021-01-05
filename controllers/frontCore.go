package controllers

import (
	"github.com/SteakBarbare/Partiel-MT4/views"
	"github.com/gin-gonic/gin"
)

var login *views.View
var form *views.View

// Load the different templates and load them in the corresponding view file
func LoginPage(c *gin.Context) {
	login = views.NewView("bootstrap", "views/login.go.html")
	login.Render(c.Writer, nil)
}

func FormPage(c *gin.Context) {
	login = views.NewView("bootstrap", "views/form.go.html")
	login.Render(c.Writer, nil)
}

// func CreateAccountPage(c *gin.Context) {
// 	contact = views.NewView("bootstrap", "views/register.go.html")
// 	contact.Render(c.Writer, nil)
// }

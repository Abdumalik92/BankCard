package pkg

import (
	"github.com/Abdumalik92/bank_card/db"
	"github.com/Abdumalik92/bank_card/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "auth.html", gin.H{})
	return
}

func Auth(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")
	db.Dbs.Where("login=? and password=?", login, password).First(&Personal)
	if Personal == (models.Personal{}) {
		Personal = models.Personal{}
		c.HTML(http.StatusNotAcceptable, "auth.html", gin.H{"proverka": "Вы ввели неправильно логин или пароль"})
		return
	}
	session := sessions.Default(c)
	session.Clear()
	session.Set("personalID", Personal.ID)
	session.Save()
	Orders(c)
	Personal = models.Personal{}
	return
}

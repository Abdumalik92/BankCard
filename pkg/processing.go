package pkg

import (
	"github.com/Abdumalik92/bank_card/db"
	"github.com/Abdumalik92/bank_card/models"
	"log"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Ordersprocessing(c *gin.Context) {
	var cardForm models.StatusCard
	c.Bind(&cardForm)

	if cardForm.Typ == "1" {
		session := sessions.Default(c)
		db.Dbs.Where("id=?", session.Get("personalID")).First(&Personal)
		var add models.Orderprocessing
		var idDepartment int
		trx := db.Dbs.Begin()
		defer func() {
			if r := recover(); r != nil {
				trx.Rollback()
			}
		}()
		Or = models.Order{}
		db.Dbs.Where("status=? AND id_personal=?", 6, Personal.ID).Find(&Or)
		if Or.IDTypeCard == 1 {
			idDepartment = 1
		} else if Or.IDTypeCard == 2 {
			idDepartment = 2
		}
		description := c.PostForm("description")

		add.Description = description
		add.Date = time.Now()
		add.IDOrder = Or.ID
		add.IDDepartment = idDepartment
		add.IDPersonal = Personal.ID
		add.Status = 2

		if err := trx.Model(&models.Order{}).Where("id=?", Or.ID).Updates(models.Order{Status: 2, Datelastupdate: time.Now(), IDPersonal: Personal.ID}).Error; err != nil {
			trx.Rollback()
			log.Println("Cann't update order " + err.Error())
			ss = "Произошла ошибка, попрбуйте познее повторить действие"
			Orders(c)
			add = models.Orderprocessing{}
			Or = models.Order{}
			Personal = models.Personal{}
			return
		}
		if err := trx.Create(&add).Error; err != nil {
			trx.Rollback()
			log.Println("Cann't create orderprocessing " + err.Error())
			ss = "Произошла ошибка, попрбуйте познее повторить действие"
			Orders(c)
			add = models.Orderprocessing{}
			Or = models.Order{}
			Personal = models.Personal{}
			return
		}
		trx.Commit()
		add = models.Orderprocessing{}
		Or = models.Order{}
		Personal = models.Personal{}

		ss = "Заявка на рассмотрении"
		Orders(c)
		ss = ""
		return
	} else if cardForm.Typ == "2" {
		session := sessions.Default(c)
		db.Dbs.Where("id=?", session.Get("personalID")).First(&Personal)
		var add models.Orderprocessing

		trx := db.Dbs.Begin()
		defer func() {
			if r := recover(); r != nil {
				trx.Rollback()
			}
		}()

		description := c.PostForm("description")
		trx.Where("status=? and id_personal=?", 6, Personal.ID).First(&Or)

		add.Description = description
		add.Date = time.Now()
		add.IDOrder = Or.ID
		add.IDDepartment = 3
		add.IDPersonal = Personal.ID
		add.Status = 5

		if err := trx.Model(&models.Order{}).Where("id=?", Or.ID).Updates(models.Order{Status: 5, Datelastupdate: time.Now(), IDPersonal: Personal.ID}).Error; err != nil {
			trx.Rollback()
			log.Println("Cann't update order " + err.Error())
			ss = "Произошла ошибка, попрбуйте познее повторить действие"
			Orders(c)
			add = models.Orderprocessing{}
			Or = models.Order{}
			Personal = models.Personal{}
			return
		}
		if err := trx.Create(&add).Error; err != nil {
			trx.Rollback()
			log.Println("Cann't create orderprocessing " + err.Error())
			ss = "Произошла ошибка, попрбуйте познее повторить действие"
			Orders(c)
			add = models.Orderprocessing{}
			Or = models.Order{}
			Personal = models.Personal{}
			return
		}
		trx.Commit()

		ss = "В заявке отказано"
		Orders(c)
		Or = models.Order{}
		Personal = models.Personal{}
		add = models.Orderprocessing{}
		ss = ""
		return
	} else {

		ss = "Вы не выбрали статус заказа"
		Orders(c)
		Or = models.Order{}
		Personal = models.Personal{}

		ss = ""

		return
	}

}

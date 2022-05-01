package pkg

import (
	"github.com/Abdumalik92/bank_card/db"
	"github.com/Abdumalik92/bank_card/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Orders(c *gin.Context) {
	session := sessions.Default(c)
	add := models.Orderprocessing{}
	Personal = models.Personal{}
	Or = models.Order{}
	db.Dbs.Where("id=?", session.Get("personalID")).First(&Personal)
	trx := db.Dbs.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()
	if Personal != (models.Personal{}) {

		if Personal.IDPosition == 3 {

			db.Dbs.Where("status=? or (status=? and id_personal=?)", 1, 6, Personal.ID).First(&Or)
			db.Dbs.Where("id=?", Or.Status).First(&Stat)
			db.Dbs.Where("id=?", Or.IDTypeCard).First(&Typs)

			if Or.PuthUploadDoc == "" {
				c.HTML(http.StatusNotFound, "orderprosseng.html", gin.H{"proverka": "Заявок нету"})
				Or = models.Order{}
				Stat = models.Status{}
				Personal = models.Personal{}
				Typs = models.Typecard{}
				add = models.Orderprocessing{}
				session.Clear()
				return
			} else {

				add = models.Orderprocessing{}
				add.Description = "Обрабатывается"
				add.Date = time.Now()
				add.IDOrder = Or.ID
				add.IDDepartment = 3
				add.IDPersonal = Personal.ID
				add.Status = 6

				if err := trx.Create(&add).Error; err != nil {
					trx.Rollback()
					log.Println("Cann't create orderprocessing " + err.Error())
					c.HTML(http.StatusInternalServerError, "auth.html", gin.H{"proverka": "Произошла ошибка, свяжитесь с администратором"})
					return
				}

				if err := trx.Model(&models.Order{}).Where("id=?", Or.ID).Updates(models.Order{Status: 6, Datelastupdate: time.Now(), IDPersonal: Personal.ID}).Error; err != nil {
					trx.Rollback()
					log.Println("Cann't update order " + err.Error())
					c.HTML(http.StatusInternalServerError, "auth.html", gin.H{"proverka": "Произошла ошибка, свяжитесь с администратором"})
					return
				}
				trx.Commit()
				c.HTML(http.StatusOK, "orderprosseng.html", gin.H{"id": Or.ID, "firstName": Or.FirstName, "lastName": Or.LastName, "tel": Or.Tel, "id_typeCard": Typs.Name, "puthUploadDoc": Files(Or.PuthUploadDoc), "date": Or.Date.Format("2006-01-02 15:04:05"), "status": Stat.Name, "proverka": ss})
				Or = models.Order{}
				Stat = models.Status{}
				Personal = models.Personal{}
				add = models.Orderprocessing{}
				Typs = models.Typecard{}
				return
			}

		} else if Personal.IDPosition == 4 {

			add = models.Orderprocessing{}

			db.Dbs.Where("status=? or (status=? and id_personal=?)", 3, 8, Personal.ID).First(&Or)
			db.Dbs.Where("id=?", Or.IDTypeCard).First(&Typs)
			db.Dbs.Where("id_order=?", Or.ID).Last(&add)

			db.Dbs.Where("id=?", Or.Status).First(&Stat)

			if Or.PuthUploadDoc == "" {

				c.HTML(200, "auth.html", gin.H{"proverka": "Заявок нету"})
				Stat = models.Status{}
				Typs = models.Typecard{}
				Or = models.Order{}
				Personal = models.Personal{}
				add = models.Orderprocessing{}
				session.Clear()
				return
			} else {
				add1 := models.Orderprocessing{}
				add1.Description = "Выпускается"
				add1.Date = time.Now()
				add1.IDOrder = Or.ID
				add1.IDDepartment = 4
				add1.IDPersonal = Personal.ID
				add1.Status = 8
				if err := trx.Create(&add1).Error; err != nil {
					trx.Rollback()
					log.Println("Cann't create orderprocessing " + err.Error())
					c.HTML(http.StatusInternalServerError, "auth.html", gin.H{"proverka": "Произошла ошибка, свяжитесь с администратором"})
					Stat = models.Status{}
					Typs = models.Typecard{}
					Or = models.Order{}
					Personal = models.Personal{}
					add = models.Orderprocessing{}
					session.Clear()
					return
				}

				if err := trx.Model(&models.Order{}).Where("id=?", Or.ID).Updates(models.Order{Status: 8, Datelastupdate: time.Now(), IDPersonal: Personal.ID}).Error; err != nil {
					trx.Rollback()
					log.Println("Cann't update order " + err.Error())
					c.HTML(http.StatusInternalServerError, "auth.html", gin.H{"proverka": "Произошла ошибка, свяжитесь с администратором"})
					Stat = models.Status{}
					Typs = models.Typecard{}
					Or = models.Order{}
					Personal = models.Personal{}
					add = models.Orderprocessing{}
					session.Clear()
					return
				}
				trx.Commit()
				c.HTML(http.StatusOK, "delevery.html", gin.H{"id": Or.ID, "firstName": Or.FirstName, "lastName": Or.LastName, "tel": Or.Tel, "puthUploadDoc": Files(Or.PuthUploadDoc), "id_typeCard": Typs.Name, "date": Or.Date.Format("2006-01-02 15:04:05"), "status": Stat.Name, "description": add.Description, "proverka": ss})
				Stat = models.Status{}
				Typs = models.Typecard{}
				Or = models.Order{}
				Personal = models.Personal{}
				add = models.Orderprocessing{}
				return
			}
		} else if Personal.IDPosition == 1 {

			db.Dbs.Where("(status=? and id_typecard=?) or (status=? and id_personal=?)", 2, 1, 7, Personal.ID).First(&Or)
			db.Dbs.Where("id=?", Or.Status).First(&Stat)
			db.Dbs.Where("id_order=?", Or.ID).Last(&add)

			if Or.PuthUploadDoc == "" {
				c.HTML(200, "auth.html", gin.H{"proverka": "Заявок нету"})
				Stat = models.Status{}
				Or = models.Order{}
				add = models.Orderprocessing{}
				Personal = models.Personal{}
				session.Clear()
				return
			} else {
				add1 := models.Orderprocessing{}
				add1.Description = "Рассмотривается"
				add1.Date = time.Now()
				add1.IDOrder = Or.ID
				add1.IDDepartment = 1
				add1.IDPersonal = Personal.ID
				add1.Status = 7
				if err := trx.Create(&add1).Error; err != nil {
					trx.Rollback()
					log.Println("Cann't create orderprocessing " + err.Error())
					c.HTML(http.StatusInternalServerError, "auth.html", gin.H{"proverka": "Произошла ошибка, свяжитесь с администратором"})
					Stat = models.Status{}
					Typs = models.Typecard{}
					Or = models.Order{}
					Personal = models.Personal{}
					add = models.Orderprocessing{}
					session.Clear()
					return
				}

				if err := trx.Model(&models.Order{}).Where("id=?", Or.ID).Updates(models.Order{Status: 7, Datelastupdate: time.Now(), IDPersonal: Personal.ID}).Error; err != nil {
					trx.Rollback()
					log.Println("Cann't update order " + err.Error())
					c.HTML(http.StatusInternalServerError, "auth.html", gin.H{"proverka": "Произошла ошибка, свяжитесь с администратором"})
					Stat = models.Status{}
					Typs = models.Typecard{}
					Or = models.Order{}
					Personal = models.Personal{}
					add = models.Orderprocessing{}
					session.Clear()
					return
				}
				trx.Commit()
				c.HTML(http.StatusOK, "debit.html", gin.H{"id": Or.ID, "firstName": Or.FirstName, "lastName": Or.LastName, "tel": Or.Tel, "puthUploadDoc": Files(Or.PuthUploadDoc), "date": Or.Date.Format("2006-01-02 15:04:05"), "status": Stat.Name, "description": add.Description, "proverka": ss})

				Stat = models.Status{}
				Or = models.Order{}
				add = models.Orderprocessing{}
				Personal = models.Personal{}
				return
			}
		} else if Personal.IDPosition == 2 {
			add = models.Orderprocessing{}
			db.Dbs.Where("(status=? and id_typecard=?) or (status=? and id_personal=?)", 2, 2, 7, Personal.ID).First(&Or)
			db.Dbs.Where("id=?", Or.Status).First(&Stat)
			db.Dbs.Where("id_order=?", Or.ID).Last(&add)

			if Or.PuthUploadDoc == "" {
				c.HTML(200, "auth.html", gin.H{"proverka": "Заявок нету"})
				Stat = models.Status{}

				add = models.Orderprocessing{}
				Or = models.Order{}
				Personal = models.Personal{}
				session.Clear()
				return
			} else {
				add1 := models.Orderprocessing{}
				add1.Description = "Рассмотривается"
				add1.Date = time.Now()
				add1.IDOrder = Or.ID
				add1.IDDepartment = 2
				add1.IDPersonal = Personal.ID
				add1.Status = 7

				if err := trx.Model(&models.Order{}).Where("id=?", Or.ID).Updates(models.Order{Status: 7, Datelastupdate: time.Now(), IDPersonal: Personal.ID}).Error; err != nil {
					trx.Rollback()
					log.Println("Cann't update order " + err.Error())
					c.HTML(http.StatusInternalServerError, "auth.html", gin.H{"proverka": "Произошла ошибка, свяжитесь с администратором"})
					Stat = models.Status{}
					Or = models.Order{}
					add = models.Orderprocessing{}
					Personal = models.Personal{}
					session.Clear()
					return
				}
				if err := trx.Create(&add1).Error; err != nil {
					trx.Rollback()
					log.Println("Cann't create orderprocessing " + err.Error())
					c.HTML(http.StatusInternalServerError, "auth.html", gin.H{"proverka": "Произошла ошибка, свяжитесь с администратором"})
					Stat = models.Status{}
					Or = models.Order{}
					add = models.Orderprocessing{}
					Personal = models.Personal{}
					session.Clear()
					return
				}
				trx.Commit()
				c.HTML(http.StatusOK, "credit.html", gin.H{"id": Or.ID, "firstName": Or.FirstName, "lastName": Or.LastName, "tel": Or.Tel, "puthUploadDoc": Files(Or.PuthUploadDoc), "date": Or.Date.Format("2006-01-02 15:04:05"), "status": Stat.Name, "description": add.Description, "proverka": ss})
				Stat = models.Status{}
				Typs = models.Typecard{}
				Or = models.Order{}
				add = models.Orderprocessing{}
				Personal = models.Personal{}

				return
			}
		}
	}
	c.Redirect(307, "/auth")
	Or = models.Order{}
	Stat = models.Status{}
	Typs = models.Typecard{}
	add = models.Orderprocessing{}
	Personal = models.Personal{}
	return
}

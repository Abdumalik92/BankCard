package pkg

import (
	"fmt"
	"github.com/Abdumalik92/bank_card/db"
	"github.com/Abdumalik92/bank_card/models"
	"log"
	"math/rand"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreditRelease(c *gin.Context) {
	session := sessions.Default(c)
	Personal = models.Personal{}
	Or = models.Order{}
	db.Dbs.Where("id=?", session.Get("personalID")).First(&Personal)
	var account models.Bankaccount
	var client models.Client
	var add models.Orderprocessing
	firtsname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	birthday := c.PostForm("birthday")
	tel := c.PostForm("tel")
	itn := c.PostForm("itn")
	adress := c.PostForm("adress")
	keyWord := c.PostForm("keyWord")
	client.Birthday = birthday
	client.Adress = adress
	client.Firstname = firtsname
	client.Lastname = lastname
	client.Tel = tel
	client.KeyWord = keyWord
	client.Itn = itn
	trs := db.Dbs.Begin()
	defer func() {
		if r := recover(); r != nil {
			trs.Rollback()
		}
	}()
	trs.Where("status=? and id_typecard=? and id_personal=?", 7, 2, Personal.ID).First(&Or)

	client.IDOrder = Or.ID

	if err := trs.Create(&client).Error; err != nil {
		trs.Rollback()
		log.Println("Cann't create client " + err.Error())
		ss = "Произошла ошибка, попрбуйте познее повторить действие"
		Orders(c)
		ss = ""
		client = models.Client{}
		Or = models.Order{}
		add = models.Orderprocessing{}
		account = models.Bankaccount{}
		Personal = models.Personal{}
		return
	}

	trs.Where("id_order=?", Or.ID).First(&client)
	account.IDClient = client.ID

	rand.Seed(time.Now().Unix())

	account.Balance = 0
	account.Number = "4081781009991" + fmt.Sprint(rand.Int63n(1e7))

	if err := trs.Create(&account).Error; err != nil {
		trs.Rollback()
		log.Println("Cann't create account " + err.Error())
		ss = "Произошла ошибка, попрбуйте познее повторить действие"
		Orders(c)
		ss = ""
		client = models.Client{}
		Or = models.Order{}
		add = models.Orderprocessing{}
		account = models.Bankaccount{}
		Personal = models.Personal{}
		return
	}

	add.IDPersonal = Personal.ID
	add.Description = "На выпуск"
	add.Date = time.Now()
	add.IDOrder = client.IDOrder
	add.IDDepartment = 4
	add.Status = 3
	if err := trs.Create(&add).Error; err != nil {
		trs.Rollback()
		log.Println("Cann't create orderprocessing " + err.Error())
		ss = "Произошла ошибка, попрбуйте познее повторить действие"
		Orders(c)
		ss = ""
		client = models.Client{}
		Or = models.Order{}
		add = models.Orderprocessing{}
		account = models.Bankaccount{}
		Personal = models.Personal{}
		return
	}

	if err := trs.Model(&models.Order{}).Where("id=?", Or.ID).Update("status", 3).Error; err != nil {
		trs.Rollback()
		log.Println("Cann't update order " + err.Error())
		ss = "Произошла ошибка, попрбуйте познее повторить действие"
		Orders(c)
		ss = ""
		client = models.Client{}
		Or = models.Order{}
		add = models.Orderprocessing{}
		account = models.Bankaccount{}
		Personal = models.Personal{}
		return
	}

	ss = "Клиент создан"
	Orders(c)
	trs.Commit()
	client = models.Client{}
	Or = models.Order{}
	add = models.Orderprocessing{}
	account = models.Bankaccount{}
	Personal = models.Personal{}
	ss = ""
	return
}

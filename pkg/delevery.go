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

func DeleveryFin(c *gin.Context) {
	session := sessions.Default(c)
	Personal = models.Personal{}
	Or = models.Order{}
	db.Dbs.Where("id=?", session.Get("personalID")).First(&Personal)
	var no models.Notification
	var card models.Card
	var typ models.Typecard
	var bankA models.Bankaccount
	var add models.Orderprocessing
	trs := db.Dbs.Begin()
	defer func() {
		if r := recover(); r != nil {
			trs.Rollback()
		}
	}()
	trs.Where("status=? and id_personal=?", 8, Personal.ID).First(&Or)
	trs.Where("id_order=?", Or.ID).First(&client)
	trs.Where("id_client=?", client.ID).First(&bankA)
	trs.Where("id=?", Or.IDTypeCard).First(&typ)

	description := c.PostForm("description")

	date := time.Now()

	card.CardOwner = client.Firstname + " " + client.Lastname
	year := fmt.Sprint(date.Year() + 3)
	card.ValidThru = fmt.Sprint(date.Format("01")) + "/" + year[len(year)-2:]
	rand.Seed(time.Now().Unix())
	card.CardNumber = fmt.Sprint(typ.PrefixNumber) + fmt.Sprint(rand.Int63n(1e10))
	card.IDTypeCard = typ.ID
	card.IDBankaccount = bankA.ID

	if err := trs.Create(&card).Error; err != nil {
		trs.Rollback()
		log.Println("Cann't create card " + err.Error())
		ss = "Произошла ошибка, попрбуйте познее повторить действие"
		Orders(c)
		ss = ""
		Or = models.Order{}
		typ = models.Typecard{}
		client = models.Client{}
		no = models.Notification{}
		card = models.Card{}
		bankA = models.Bankaccount{}
		add = models.Orderprocessing{}
		Personal = models.Personal{}
		return
	}

	no.Tel = client.Tel
	no.Date = time.Now()
	no.IDOrder = Or.ID
	no.Message = "Ваша карта готово"
	if err := trs.Create(&no).Error; err != nil {
		trs.Rollback()
		log.Println("Cann't create notification " + err.Error())
		ss = "Произошла ошибка, попрбуйте познее повторить действие"
		Orders(c)
		ss = ""
		Or = models.Order{}
		typ = models.Typecard{}
		client = models.Client{}
		no = models.Notification{}
		card = models.Card{}
		bankA = models.Bankaccount{}
		add = models.Orderprocessing{}
		Personal = models.Personal{}
		return
	}
	add.Description = description
	add.Date = time.Now()
	add.IDOrder = client.IDOrder
	add.IDDepartment = 3
	add.Status = 4
	add.IDPersonal = Personal.ID
	if err := trs.Create(&add).Error; err != nil {
		trs.Rollback()
		log.Println("Cann't create orderprocessing " + err.Error())
		ss = "Произошла ошибка, попрбуйте познее повторить действие"
		Orders(c)
		ss = ""
		Or = models.Order{}
		typ = models.Typecard{}
		client = models.Client{}
		no = models.Notification{}
		card = models.Card{}
		bankA = models.Bankaccount{}
		add = models.Orderprocessing{}
		Personal = models.Personal{}
		return
	}
	if err := trs.Model(&models.Order{}).Where("id=?", Or.ID).Update("status", 4).Error; err != nil {
		trs.Rollback()
		log.Println("Cann't update order " + err.Error())
		ss = "Произошла ошибка, попрбуйте познее повторить действие"
		Orders(c)
		ss = ""
		Or = models.Order{}
		typ = models.Typecard{}
		client = models.Client{}
		no = models.Notification{}
		card = models.Card{}
		bankA = models.Bankaccount{}
		add = models.Orderprocessing{}
		Personal = models.Personal{}
		return
	}
	trs.Commit()

	ss = "Карта создана"
	Orders(c)
	Or = models.Order{}
	typ = models.Typecard{}
	client = models.Client{}
	no = models.Notification{}
	card = models.Card{}
	bankA = models.Bankaccount{}
	add = models.Orderprocessing{}
	Personal = models.Personal{}
	ss = ""
	return
}

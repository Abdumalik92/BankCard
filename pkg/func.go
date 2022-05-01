package pkg

import (
	"fmt"
	"github.com/Abdumalik92/bank_card/db"
	"github.com/Abdumalik92/bank_card/models"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var client models.Client
var Or models.Order

var Account models.Bankaccount
var Stat models.Status
var Typs models.Typecard
var Dep models.Department
var Personal models.Personal

var ss string

func MainPage(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{})
	return
}

func CardOrder(c *gin.Context) {

	var cardForm models.MyCard
	var tCard int
	var i int
	c.Bind(&cardForm)
	rand.Seed(time.Now().Unix())
	i = rand.Intn(1000000)
	firtsname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	tel := c.PostForm("tel")
	form, _ := c.MultipartForm()
	if cardForm.Typecar == "1" {
		tCard = 1

	} else if cardForm.Typecar == "2" {
		tCard = 2
	} else {

		c.HTML(http.StatusOK, "index.html", gin.H{"proverka": "Вы не выбрали тип карты"})

		return
	}

	t := time.Now().Format("2006_01_02")

	puthUploadDoc := "upload/" + t + fmt.Sprint(i) + "_" + lastname + "/"

	Or.IDTypeCard = tCard
	Or.Date = time.Now()
	Or.FirstName = firtsname
	Or.LastName = lastname
	Or.Tel = tel
	Or.Status = 1
	Or.PuthUploadDoc = puthUploadDoc
	Or.Datelastupdate = time.Now()
	Or.IDPersonal = 8
	trs := db.Dbs.Begin()
	defer func() {
		if r := recover(); r != nil {
			trs.Rollback()
		}
	}()

	if err := trs.Create(&Or).Error; err != nil {
		trs.Rollback()
		log.Println("Cann't create order " + err.Error())
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{"proverka": "Произошла ошибка, попробуйте позднее"})
		return
	}
	if _, err := os.Stat("upload/" + t + fmt.Sprint(i) + "_" + lastname); os.IsNotExist(err) {
		os.Mkdir("upload/"+t+fmt.Sprint(i)+"_"+lastname, 0522)
	}

	files := form.File["files"]
	for _, file := range files {
		err := c.SaveUploadedFile(file, "upload/"+t+fmt.Sprint(i)+"_"+lastname+"/"+t+fmt.Sprint(i)+file.Filename)
		if err != nil {
			trs.Rollback()
			log.Println(err.Error())
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{"proverka": "Произошла ошибка, попробуйте позднее"})
			Or = models.Order{}
			return
		}
	}

	trs.Commit()

	c.HTML(http.StatusOK, "index.html", gin.H{"proverka": "Ваша заявка принято"})
	Or = models.Order{}
	return
}

func Files(puth string) []string {
	var d []string

	dir, err := os.Open(puth)
	if err != nil {
		log.Println("Cann't open dir ", err.Error)
		return d
	}
	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		log.Println("Cann't read into dir ", err.Error)
		return d
	}
	for _, fi := range fileInfos {

		d = append(d, puth+fi.Name())

	}
	return d
}

func UpdateStatus() {
	for {
		timeout := time.After(10 * time.Second)
		select {
		case <-timeout:
			db.Dbs.Exec("call updateorders()")
		}
	}
}

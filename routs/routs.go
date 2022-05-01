package routs

import (
	"fmt"
	"github.com/Abdumalik92/bank_card/db"
	"github.com/Abdumalik92/bank_card/logs"
	"github.com/Abdumalik92/bank_card/pkg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func Init() {
	r := gin.Default()
	f, _ := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE, 0666)
	gin.DefaultWriter = io.MultiWriter(f)
	logger := logrus.New()
	logger.Level = logrus.TraceLevel
	logger.SetOutput(gin.DefaultWriter)
	r.Use(logs.Logger(logger), gin.Recovery())
	db.Dbs = db.Open(logger)
	defer db.Dbs.Close()

	store := cookie.NewStore([]byte("secret"))
	//store.Options = sessions.Option{Path: "/auth", MaxAge: 860000, Secure: true}
	r.Use(sessions.Sessions("sessionID", store))
	r.LoadHTMLGlob("tmp/*")
	r.GET("/", pkg.MainPage)
	r.POST("/", pkg.CardOrder)
	r.GET("/auth", pkg.Login)
	r.POST("/auth", pkg.Auth)
	r.GET("/order", pkg.Orders)
	r.POST("/underwrite", pkg.Ordersprocessing)
	r.POST("/credit", pkg.CreditRelease)
	r.POST("/debit", pkg.DebitRelease)
	r.POST("/delevery", pkg.DeleveryFin)
	r.Static("/css", "./css")
	r.Static("/fonts", "./fonts")
	r.Static("/img", "./img")
	r.Static("/script", "./script")
	r.Static("/upload", "./upload")

	fmt.Println("Server is listening...")
	r.Run(":8080")
}

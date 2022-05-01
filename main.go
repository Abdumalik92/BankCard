package main

import (
	"github.com/Abdumalik92/bank_card/pkg"
	"github.com/Abdumalik92/bank_card/routs"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	routs.Init()
	go pkg.UpdateStatus()
}

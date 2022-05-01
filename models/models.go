package models

import "time"

type Order struct {
	ID             int       `gorm:"column:id;primary_key;NOT NULL;"`
	FirstName      string    `gorm:"column:firstname;size:20;NOT NULL;"`
	LastName       string    `gorm:"column:lastname;size:20;NOT NULL;"`
	Tel            string    `gorm:"column:tel;size:15;NOT NULL;"`
	IDTypeCard     int       `gorm:"column:id_typecard;NOT NULL;"`
	PuthUploadDoc  string    `gorm:"column:puthuploaddoc;size:60;NOT NULL;"`
	Date           time.Time `gorm:"column:date;size:20;NOT NULL;"`
	Status         int       `gorm:"column:status;NOT NULL;"`
	Datelastupdate time.Time `gorm:"column:date_last_update;NOT NULL;"`
	IDPersonal     int       `gorm:"column:id_personal;NOT NULL;"`
}
type Card struct {
	ID            int    `gorm:"column:id;primary_key;NOT NULL;"`
	CardNumber    string `gorm:"column:cardNumber;size:16;NOT NULL;unique;"`
	CardOwner     string `gorm:"column:cardOwner;size:60;NOT NULL;"`
	ValidThru     string `gorm:"column:validThru;size:5;NOT NULL;"`
	IDTypeCard    int    `gorm:"column:id_typecard;NOT NULL;"`
	IDBankaccount int    `gorm:"column:id_bankaccount;NOT NULL;"`
}

type Client struct {
	ID        int    `gorm:"column:id;primary_key;NOT NULL;"`
	Firstname string `gorm:"column:firstname;size:20;NOT NULL;"`
	Lastname  string `gorm:"column:lastname;size:20;NOT NULL;"`
	Birthday  string `gorm:"column:birthday;size:15;NOT NULL;"`
	Tel       string `gorm:"column:tel;size:20;NOT NULL;"`
	Itn       string `gorm:"column:itn;size:12;NOT NULL;unique;"`
	Adress    string `gorm:"column:adress;size:90;NOT NULL;"`
	IDOrder   int    `gorm:"column:id_order;NOT NULL;unique;"`
	KeyWord   string `gorm:"column:keyWord;size:45;NOT NULL;"`
}

type Department struct {
	ID   int    `gorm:"column:id;primary_key;NOT NULL;"`
	Name string `gorm:"column:name;size:30;NOT NULL;"`
}
type Status struct {
	ID   int    `gorm:"column:id;primary_key;NOT NULL;"`
	Name string `gorm:"column:name;size:20;NOT NULL;"`
}

type Orderprocessing struct {
	ID           int       `gorm:"column:id;primary_key;NOT NULL;"`
	IDOrder      int       `gorm:"column:id_order;NOT NULL;"`
	IDDepartment int       `gorm:"column:id_department;NOT NULL;"`
	IDPersonal   int       `gorm:"column:id_personal;NOT NULL;"`
	Date         time.Time `gorm:"column:date;NOT NULL;"`
	Status       int       `gorm:"column:status;NOT NULL;"`
	Description  string    `gorm:"column:description;size:100;NOT NULL;"`
}

type Personal struct {
	ID         int    `gorm:"column:id;primary_key;NOT NULL;"`
	Firstname  string `gorm:"column:firstname;size:20;NOT NULL;"`
	Lastname   string `gorm:"column:lastname;size:20;NOT NULL;"`
	Tel        string `gorm:"column:tel;size:15;NOT NULL;"`
	Adress     string `gorm:"column:adress;size:90;NOT NULL;"`
	IDPosition int    `gorm:"column:id_position;NOT NULL;"`
	Login      string `gorm:"column:login;size:20;NOT NULL;unique;"`
	Password   string `gorm:"column:password;size:20;NOT NULL;"`
}
type Position struct {
	ID           int    `gorm:"column:id;primary_key;NOT NULL;"`
	Name         string `gorm:"column:name;size:30;NOT NULL;"`
	IDDepartment int    `gorm:"column:id_department;NOT NULL;"`
}
type Typecard struct {
	ID           int    `gorm:"column:id;primary_key;NOT NULL;"`
	Name         string `gorm:"column:name;size:30;NOT NULL;"`
	PrefixNumber int    `gorm:"column:prefixNumber;size:6;NOT NULL;"`
	CardColor    string `gorm:"column:cardColor;size:15;NOT NULL;"`
}
type Notification struct {
	ID      int       `gorm:"column:id;primary_key;NOT NULL;"`
	IDOrder int       `gorm:"column:id_order;NOT NULL;unique;"`
	Message string    `gorm:"column:message;size:90;NOT NULL;"`
	Date    time.Time `gorm:"column:date;size:15;NOT NULL;"`
	Tel     string    `gorm:"column:tel;size:20;NOT NULL;"`
}
type MyCard struct {
	Typecar string `form:"list1"`
}
type StatusCard struct {
	Typ string `form:"list2"`
}

type Bankaccount struct {
	ID       int     `gorm:"column:id;primary_key;NOT NULL;"`
	Number   string  `gorm:"column:number;size:20;NOT NULL;unique;"`
	Balance  float64 `gorm:"column:balance;NOT NULL;"`
	IDClient int     `gorm:"column:id_client;NOT NULL;"`
}

type Respons struct {
	ID            int       `json:"id,omitempty"`
	IDOrder       int       `json:"id_order,omitempty"`
	IDPosition    int       `json:"id_position,omitempty"`
	IDPersonal    int       `json:"id_personal,omitempty"`
	IDCard        int       `json:"id_card,omitempty"`
	IDTypeCard    int       `json:"id_typeCard,omitempty"`
	IDDepartment  int       `json:"id_department,omitempty"`
	FirstName     string    `json:"firstname,omitempty"`
	LastName      string    `json:"lastname,omitempty"`
	Birthday      string    `json:"birthday,omitempty"`
	Tel           string    `json:"tel,omitempty"`
	Adress        string    `json:"adress,omitempty"`
	Itn           int       `json:"itn,omitempty"`
	IDBankaccount int       `json:"bankaccount,omitempty"`
	KeyWord       string    `json:"keyWord,omitempty"`
	Name          string    `json:"name,omitempty"`
	Balance       float64   `json:"balance,omitempty"`
	CardNumber    int       `json:"cardNumber,omitempty"`
	CardOwner     string    `json:"cardOwner,omitempty"`
	ValidThru     string    `json:"validThru,omitempty"`
	PuthUploadDoc string    `json:"puthuploaddoc,omitempty"`
	PrefixNumber  int       `json:"prefixNumber,omitempty"`
	CardColor     string    `json:"cardColor,omitempty"`
	Date          time.Time `json:"date,omitempty"`
	Status        string    `json:"status,omitempty"`
	Message       string    `json:"message,omitempty"`
	Description   string    `json:"description,omitempty"`
}
type Resp struct {
	ID            int
	Firstname     string
	Lastname      string
	Tel           string
	Puthuploaddoc string
	Date          time.Time
	Id_typecard   int
	Status        int
	Description   string
}

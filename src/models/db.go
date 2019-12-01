package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

var ormObject orm.Ormer

func ConnectToDb() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default","postgres", "user=postgres password=docker dbname=postgres host=localhost port=5432 sslmode=disable")
	orm.RegisterModel(new(Users))
	ormObject = orm.NewOrm()
}

func GetOrmObject() orm.Ormer {
	return ormObject
}
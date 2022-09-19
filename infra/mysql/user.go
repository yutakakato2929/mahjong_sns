package mysql

import (
	"fmt"
	"mahjong_sns/infra/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account  string `db:"account" form:"account" binding:"required"`
	Name     string `db:"name" form:"name" binding:"required"`
	Password string `db:"password" form:"password" binding:"required"`
}

func dbConnect() (*gorm.DB, error) {
	confDB, err := conf.ReadConfDB()
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true", confDB.User, confDB.Pass, confDB.Host, confDB.Port, confDB.DbName, confDB.Charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func DbInit() {
	db, err := dbConnect()
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&User{})
}

func UserInsert(acc, name, pw string) {
	db, _ := dbConnect()
	db.Create(&User{Account: acc, Name: name, Password: pw})
}

func UserUpdate(id int, acc, name, pw string) {
	db, _ := dbConnect()
	var user User
	db.First(&user, id)
	user.Account = acc
	user.Name = name
	user.Password = pw
	db.Save(&user)
}

func UserGetAll() []User {
	db, _ := dbConnect()
	var users []User
	db.Order("created_at").Find(&users)
	return users
}

func UserGetOne(id int) User {
	db, _ := dbConnect()
	var user User
	db.First(&user, id)
	return user
}

func UserDelete(id int) {
	db, _ := dbConnect()
	db.Where("id = ?", id).Delete(&User{})
}

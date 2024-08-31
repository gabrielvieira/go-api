package model

import (
	"gorm.io/gorm"
)

/*
	query docs - https://gorm.io/docs/query.html
	query multiple users - db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users)
	query single user -  db.Where(&model.User{Name: "d"}).First(&user)
	selecting fields - d.Select([]string{"name", "password"}).Where(&model.User{Name: "Gabriel"}).First(&user)

	.first - returns a single entry
	.find - returns a list of entries

	handle errors

	result.RowsAffected // returns found records count, equals `len(users)`
	result.Error        // returns error
*/

type User struct {
	gorm.Model
	Name     string
	Password string
}

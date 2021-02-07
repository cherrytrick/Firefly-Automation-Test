package models

import (
	"worker/database"
)

type Test struct {
	Id      int
	Testcol string `gorm:"column:testcol"`
}

// 设置Test的表名为`test`
// func (Test) TableName() string {
//     return "test"
// }

func (this *Test) Insert() (id int, err error) {
	result := database.DB.Create(&this)
	id = this.Id
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

package service

import (
	"worker/models"
)

type Test struct {
	Id      int    `json:"id"`
	Testcol string `json:"testcol"`
}

func (this *Test) Insert() (id int, err error) {
	var testModel models.Test
	testModel.Id = this.Id
	testModel.Testcol = this.Testcol
	id, err = testModel.Insert()
	return
}

type Script struct {
	Name string `json:"name"`
}

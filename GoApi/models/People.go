package models

import (
	"fmt"
)

type Person struct {
	//gorm.Model
	Id uint `json:"id"`
	Name string `json:"Name"`
	Age uint `json:"Age"`
}

func GetPeople() []*Person {
	people := make([]*Person, 0)
	err := GetDb().Table("people").Find(&people).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return people
}
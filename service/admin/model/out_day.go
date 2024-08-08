// Code generated by http://www.gotool.top
package model

import "time"

type OutDay struct {
	Date  time.Time `gorm:"column:date;primary_key;NOT NULL" json:"date"`
	Money float32   `gorm:"column:money;default:NULL;comment:'时间'" json:"money"`
}

func (o *OutDay) TableName() string {
	return "out_day"
}
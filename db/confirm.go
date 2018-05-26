package db

import (
	"time"
)

//Confirm 粉丝参加活动结账后帮其证实真实有效
type Confirm struct {
	ID        uint       `gorm:"primary_key"`
	FansID    uint       `gorm:"index:id"` //粉丝ID
	TaskID    uint       `gorm:"index:id"` //活动ID
	DeletedAt *time.Time `sql:"index"`
}

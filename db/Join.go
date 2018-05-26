package db

import (
	"time"
)

//Join 粉丝参加活动成绩
type Join struct {
	ID           uint      `gorm:"primary_key"`
	Fore         uint      `gorm:"index:id"` //有效火力
	FansID       uint      `gorm:"index:id"` //粉丝ID
	TaskID       uint      `gorm:"index:id"` //活动ID
	ReachAt      time.Time //达成时间
	SettlementAt time.Time //结算时间
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `sql:"index"`
}

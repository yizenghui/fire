package db

import (
	"time"
)

//Report 粉丝举报活动
type Report struct { // 举报需知： 受理存在以下情况的活动，虚假、挂羊头卖狗肉、额外收费或条件直接影响活动结算的
	ID        uint   `gorm:"primary_key"`
	FansID    uint   `gorm:"index:id"`            //粉丝ID
	TaskID    uint   `gorm:"index:id"`            //活动ID
	Intro     string `gorm:"type:varchar(1024);"` //描述
	Images    string `gorm:"type:text;"`          //证图
	State     int16  //待处理 已撤消 受理中 协商解决 实锤 假锤
	Reply     string `gorm:"type:varchar(1024);"` //答复
	Contact   string //留下联系方式 微信号或者手机号
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

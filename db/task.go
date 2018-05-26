package db

import (
	"time"
)

// Task 活动任务
type Task struct {
	ID     uint   `gorm:"primary_key"`
	FansID uint   //发起人id
	City   string `sql:"index"`                // 城市(发起人所在的)
	Title  string `gorm:"type:varchar(64);"`   // 活动标题
	Intro  string `gorm:"type:varchar(1024);"` // 活动描述
	// Statement        string    `gorm:"type:varchar(1024);"` // 声明
	// TotalNum         int64     //总访问量
	// Number           int64     //最大可获奖人数
	// CompletionNumber int64     //当前完成人数
	// Fore             int64     //最低获取火力条件
	// StartAt          time.Time //开始时间
	// EndAt            time.Time //结束时间
	// Images           string    `gorm:"type:text;"` // 展品图片
	CreatedAt time.Time
	UpdatedAt time.Time
	// SpreadAt         *time.Time `sql:"index:date"` //推广期截止时间
	// ModeratedAt      *time.Time `sql:"index:date"` //审核时间
	// DeletedAt        *time.Time `sql:"index:date"`
}

// GetTaskByID Task 如果没有的话进行初始化
func (task *Task) GetTaskByID(id int64) {
	DB().First(&task, id)
}

// GetTaskByCity 获取当前城市通过审核的推广活动
func (task *Task) GetTaskByCity(city string) (tasks []Task) {
	DB().Where("moderated_at < ? and spread_at > ?", time.Now, time.Now).Where(&Task{City: city}).Find(&tasks)
	return
}

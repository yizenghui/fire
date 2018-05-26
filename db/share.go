package db

import (
	"time"
)

//Share 分享记录
type Share struct {
	ID        uint `gorm:"primary_key"`
	FansID    uint `gorm:"index:index"`
	TaskID    uint `gorm:"index:index"`
	CreatedAt time.Time
}

// Save Share
func (share *Share) Save() {
	DB().Save(&share)
}

// Log 记录 openID 分享 Task 记录
func (share *Share) Log(openID string, taskID int) {
	var fans = Fans{}
	fans.GetFansByOpenID(openID)
	share.FansID = fans.ID
	share.TaskID = uint(taskID)
	DB().Create(share)
}

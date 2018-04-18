package minappapi

import (
	// "fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Fans 粉丝
type Fans struct {
	ID         uint   `gorm:"primary_key"`
	OpenID     string `gorm:"type:varchar(255);unique_index"`
	UnionID    string
	NickName   string
	Gender     int
	City       string
	Province   string
	Country    string
	AvatarURL  string
	Language   string
	Timestamp  int64
	AppID      string
	SessionKey string // 粉丝上次的session key 如果有变化，同步一次粉丝数据
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}

//Share 分享记录
type Share struct {
	ID        uint `gorm:"primary_key"`
	FansID    uint `gorm:"index"`
	PostID    uint
	SubNum    int64 // 订阅次数 用户每提交一次+1
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Activity 提交的url
type Activity struct {
	ID        uint      `gorm:"primary_key"`
	Title     string    `gorm:"type:varchar(64);"`   // 活动标题
	Intro     string    `gorm:"type:varchar(1024);"` // 活动描述
	TotalNum  int64     //总访问量
	Number    int64     //最大可获奖人数
	Fore      int64     //最低获取火力条件
	StartAt   int64     //开始时间
	EndAt     int64     //结束时间
	Images    string    `gorm:"type:text;"` // 展品图片
	SpreadAt  time.Time `sql:"index"`       //推广期
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

//Join 粉丝参加活动成绩
type Join struct {
	ID           uint      `gorm:"primary_key"`
	FansID       uint      `gorm:"index"` //粉丝ID
	ActivityID   uint      `gorm:"index"` //活动ID
	ReachAt      time.Time //达成时间
	SettlementAt time.Time //结算时间
	Fore         int64     //有效火力
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `sql:"index"`
}

//Push 助力记录
type Push struct {
	ID        uint `gorm:"primary_key"`
	FansID    uint `gorm:"index:user_id"` // 谁在收集助力
	JoinID    uint `gorm:"index:user_id"` // 参加活动凭证
	PushID    uint `gorm:"index:user_id"` // 哪个朋友来给助力
	CreatedAt time.Time
}

// Feedback 粉丝
type Feedback struct {
	ID        uint   `gorm:"primary_key"`
	OpenID    string `gorm:"type:varchar(255);index"` // 微信文章地址
	FormID    string `gorm:"type:varchar(255);"`      //订阅formID，一次订阅只能推送一次通知
	Problem   string `gorm:"type:text;"`              // 问题
	Answer    string `gorm:"type:text;"`              // 答复
	Show      bool   //是否显示
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Post 提交的url
type Post struct {
	ID               uint   `gorm:"primary_key"`
	Title            string `gorm:"type:varchar(1024);"`             // 微信文章地址
	URL              string `gorm:"type:varchar(1024);unique_index"` // 微信文章地址
	SubNum           int64  // 订阅人次 用户每提交一次+1
	FolNum           int64  // 当前关注人数 注，如果有人关注，每过8小时检查更新
	ShareNum         int64  `gorm:"index"`      // 分享次数 现在用于排序
	ChapterFragments string `gorm:"type:text;"` // 章节片段
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time `sql:"index"`
}

// Subscribe 粉丝
type Subscribe struct {
	ID        uint   `gorm:"primary_key"`
	FansID    uint   `sql:"index"`               //粉丝 ID
	OpenID    string `gorm:"type:varchar(255);"` //提交的openid
	PostID    uint   `sql:"index"`               //post ID
	FormID    string `gorm:"type:varchar(255);"` //订阅formID，一次订阅只能推送一次通知
	Push      bool   //是否推送
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"` //删除后不再推送
}

var db *gorm.DB

//DB 返回 *gorm.DB
func DB() *gorm.DB {
	if db == nil {

		newDb, err := newDB()
		if err != nil {
			panic(err)
		}
		newDb.DB().SetMaxIdleConns(10)
		newDb.DB().SetMaxOpenConns(100)

		newDb.LogMode(false)
		db = newDb
	}

	return db
}

func newDB() (*gorm.DB, error) {

	// sqlConnection := fmt.Sprintf(
	// 	"host=%v user=%v port=%v dbname=%v sslmode=%v password=%v",
	// 	config.Database.Host,
	// 	config.Database.User,
	// 	config.Database.Port,
	// 	config.Database.Dbname,
	// 	config.Database.Sslmode,
	// 	config.Database.Password,
	// )
	// db, err := gorm.Open(config.Database.Type, sqlConnection)
	db, err := gorm.Open("sqlite3", "fireapi.db")

	if err != nil {
		return nil, err
	}
	return db, nil
}

package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	cpi "github.com/yizenghui/fire"
)

// Task 提交的url
type Task struct {
	ID               uint   `gorm:"primary_key"`
	FansID           uint   //发起人id
	City             string `sql:"index"`                // 城市(发起人所在的)
	Title            string `gorm:"type:varchar(64);"`   // 活动标题
	Intro            string `gorm:"type:varchar(1024);"` // 活动描述
	TotalNum         int64  //总访问量
	Number           int64  //最大可获奖人数
	CompletionNumber int64  //当前完成人数
	Fore             int64  //最低获取火力条件
	StartAt          int64  //开始时间
	EndAt            int64  //结束时间
	Images           string `gorm:"type:text;"` // 展品图片
	CreatedAt        time.Time
	UpdatedAt        time.Time
	SpreadAt         *time.Time `sql:"index:date"` //推广期截止时间
	ModeratedAt      *time.Time `sql:"index:date"` //审核时间
	DeletedAt        *time.Time `sql:"index:date"`
}

// NewTask 创建新任务 post
// /newstask post
func NewTask(c echo.Context) error {

	fans, e := getUser(getOpenID(c))
	if e != nil {
		return echo.ErrUnauthorized
	}

	title := c.FormValue("title")
	intro := c.FormValue("intro")
	if title != `` {

		t := cpi.Task{
			FansID: fans.ID,
			City:   fans.City,
			Title:  title,
			Intro:  intro,
		}
		cpi.DB().Create(&t)

		return c.JSON(http.StatusOK, t)
	}
	return echo.ErrUnauthorized
}

// GetTaskInfo 获取一个活动详细
// taskinfo/:id  get
func GetTaskInfo(c echo.Context) error {
	// fans, e := getUser(getOpenID(c))
	// if e != nil {
	// 	return echo.ErrUnauthorized
	// }
	id, _ := strconv.Atoi(c.Param("id"))
	// id, _ := strconv.Atoi(c.QueryParam("id"))
	var t = cpi.Task{}
	t.GetTaskByID(int64(id))

	return c.JSON(http.StatusOK, t)
}

// JoinTask 加入一个活动
// jointask/:id  get
func JoinTask(c echo.Context) error {
	fans, e := getUser(getOpenID(c)) // 获取fans信息
	if e != nil {
		return echo.ErrUnauthorized
	}
	id, _ := strconv.Atoi(c.Param("id")) //获取task id
	var t = cpi.Task{}
	t.GetTaskByID(int64(id)) // 获取task t

	nowdate := time.Now() // 当前时间

	if t.ID > 0 && nowdate.After(t.StartAt) && nowdate.Before(t.EndAt) { // 有id 后于开始时间 前于结束时间 (活动期限内)
		// if t.StartAt
		var join = cpi.Join{}
		cpi.DB().Where(&cpi.Join{TaskID: uint(id), FansID: fans.ID}).First(join)
		if join.ID == 0 {
			// 新增
			newJoin := cpi.Join{
				TaskID: uint(id),
				FansID: fans.ID,
			}
			cpi.DB().Create(&newJoin)
			return c.JSON(http.StatusOK, newJoin)
		}
	}
	return c.JSON(http.StatusOK, t)
}

// CheckJoin 检查能否加入一个活动
// checkjoin/:id  get
func CheckJoin(c echo.Context) error {
	fans, e := getUser(getOpenID(c)) // 获取fans信息
	if e != nil {
		return echo.ErrUnauthorized
	}
	id, _ := strconv.Atoi(c.Param("id")) //获取task id
	var t = cpi.Task{}
	t.GetTaskByID(int64(id)) // 获取task t

	nowdate := time.Now() // 当前时间

	if t.ID > 0 && nowdate.After(t.StartAt) && nowdate.Before(t.EndAt) { // 有id 后于开始时间 前于结束时间 (活动期限内)
		// if t.StartAt
		var join = cpi.Join{}
		cpi.DB().Where(&cpi.Join{TaskID: uint(id), FansID: fans.ID}).First(join)
		if join.ID == 0 {
			// 新增
			newJoin := cpi.Join{
				TaskID: uint(id),
				FansID: fans.ID,
			}
			cpi.DB().Create(&newJoin)
			return c.JSON(http.StatusOK, newJoin)
		}
		return c.JSON(http.StatusOK, join)
	}
	return c.JSON(http.StatusOK, t)
}

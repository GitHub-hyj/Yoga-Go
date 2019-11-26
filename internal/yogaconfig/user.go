package yogaconfig

import "time"

//Mysql 5.5只支持一个CURRENT_TIMESTAMP的默认值，因此时间的默认值都使用蓝眼云盘第一个发布版本时间 2018-01-01 00:00:00
type Base struct {
	Uuid       string    `json:"uuid"`
	Sort       int64     `json:"sort"`
	UpdateTime time.Time `json:"updateTime"`
	CreateTime time.Time `json:"createTime"`
}


type User struct {
	Base
	Role      string    `json:"role"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Gender    string    `json:"gender"`
	City      string    `json:"city"`
	AvatarUrl string    `json:"avatarUrl"`
	LastIp    string    `json:"lastIp"`
	LastTime  time.Time `json:"lastTime"`
	SizeLimit int64     `json:"sizeLimit"`
	Status    string    `json:"status"`
}

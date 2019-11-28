package yogaconfig

import (
	"errors"
	"time"
)


var (
	//ErrNotLogin 未登录帐号错误
	ErrNotLogin = errors.New("user not login")
	//ErrConfigFilePathNotSet 未设置配置文件
	ErrConfigFilePathNotSet = errors.New("config file not set")
	//ErrConfigFileNotExist 未设置Config, 未初始化
	ErrConfigFileNotExist = errors.New("config file not exist")
	//ErrConfigFileNoPermission Config文件无权限访问
	ErrConfigFileNoPermission = errors.New("config file permission denied")
	//ErrConfigContentsParseError 解析Config数据错误
	ErrConfigContentsParseError = errors.New("config contents parse error")
)


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

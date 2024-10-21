package model

import "gorm.io/gorm"

type Deploy struct {
	gorm.Model
	JobName   string `gorm:"column:name;type:varchar(20);not null;comment:'项目名称'" json:"jobName"`
	GitUrl    string `gorm:"column:giturl;type:varchar(255);comment:'git地址'" json:"gitUrl"`
	Branch    string `gorm:"type:varchar(255);comment:'分支'" json:"branch"`
	RunConfig string `gorm:"column:runconfig;type:varchar(255);comment:'启动参数'" json:"runConfig"`
}

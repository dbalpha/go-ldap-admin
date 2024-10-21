package model

import (
	"gorm.io/gorm"
	"time"
)

/**
  @author: alpha
  @since: 2024/10/11
**/

type DeployHistory struct {
	gorm.Model
	Deploy     *Deploy   `gorm:"foreigner:DeployID"`
	DeployID   uint      `gorm:"column:deploy_id;comment:'DeployID'"`
	DeployName string    `gorm:"column:deploy_name;comment:'DeployName';not null"`
	State      int       `gorm:"column:state;type:int;default:0;not null;comment:'发布状态'"`
	Env        int       `gorm:"column:env;type:int;default:0;not null;comment:'发布环境'"`
	StartAt    time.Time `gorm:"column:start_at;not null;comment:'发布开始时间'"`
	EndAt      time.Time `gorm:"column:end_at;not null;comment:'发布结束时间'"`
}

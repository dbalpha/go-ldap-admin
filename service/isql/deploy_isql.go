package isql

import (
	"errors"
	"fmt"
	"github.com/dbalpha/go-ldap-admin/model"
	"github.com/dbalpha/go-ldap-admin/model/request"
	"github.com/dbalpha/go-ldap-admin/public/common"
	"github.com/dbalpha/go-ldap-admin/public/tools"
	"gorm.io/gorm"
)

type DeployService struct{}

// Exist 判断是否存在
func (s DeployService) Exist(filter map[string]interface{}) bool {
	var dataObj model.Deploy
	err := common.DB.Debug().Order("created_at DESC").Where(filter).First(&dataObj).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// Add 添加发布
func (s DeployService) Add(deploy *model.Deploy) error {
	return common.DB.Model(&model.Deploy{}).Create(deploy).Error
}

// Delete 删除发布
func (s DeployService) Delete(deployIds []uint) error {
	return common.DB.Where("id IN (?)", deployIds).Unscoped().Delete(&model.Deploy{}).Error
}

// Update 更新发布
func (s DeployService) Update(deploy *model.Deploy) error {
	return common.DB.Model(&model.Deploy{}).Where("id = ?", deploy.ID).Updates(deploy).Error
}

// Count 获取所有job数目
func (s DeployService) Count() (int64, error) {
	var count int64
	err := common.DB.Model(&model.Deploy{}).Count(&count).Error
	return count, err
}

// Find 获取详细信息
func (s DeployService) Find(filter map[string]interface{}, data *model.Deploy) error {
	return common.DB.Where(filter).First(data).Error
}

// List 获取列表
func (s DeployService) List(req *request.DeployListReq) (deploys []*model.Deploy, err error) {
	db := common.DB.Order("name")
	jobName := req.JobName
	if jobName != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", jobName))
	}
	GitUrl := req.GitUrl
	if GitUrl != "" {
		db = db.Where("giturl LIKE ?", fmt.Sprintf("%%%s%%", GitUrl))
	}
	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err = db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Find(&deploys).Debug().Error
	return deploys, err
}

// ListCount 获取符合条件的列表数量
func (s DeployService) ListCount(req *request.DeployListReq) (count int64, err error) {
	db := common.DB.Model(&model.Deploy{})
	jobName := req.JobName
	if jobName != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", jobName))
	}
	GitUrl := req.GitUrl
	if GitUrl != "" {
		db = db.Where("giturl LIKE ?", fmt.Sprintf("%%%s%%", GitUrl))
	}
	err = db.Count(&count).Error
	return count, err
}

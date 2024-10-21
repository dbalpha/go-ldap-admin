package isql

import (
	"fmt"
	"github.com/dbalpha/go-ldap-admin/model"
	"github.com/dbalpha/go-ldap-admin/model/request"
	"github.com/dbalpha/go-ldap-admin/public/common"
	"github.com/dbalpha/go-ldap-admin/public/tools"
)

/**
  @author: alpha
  @since: 2024/10/11
**/

type DeployHisService struct{}

// Add 添加发布历史
func (receiver DeployHisService) Add(history *model.DeployHistory) error {
	err := common.DB.AutoMigrate(&model.DeployHistory{})
	if err != nil {
		return err
	}
	return common.DB.Model(&model.DeployHistory{}).Create(&history).Error
}

// Delete 删除历史记录
func (receiver DeployHisService) Delete(historyId []uint) error {
	return common.DB.Where("id in (?)", historyId).Unscoped().Delete(model.DeployHistory{}).Error
}

// Find 获取详细信息
func (receiver DeployHisService) Find(filter map[string]interface{}, data *model.DeployHistory) error {
	return common.DB.Where(filter).First(data).Error
}

// List 获取列表
func (receiver DeployHisService) List(req *request.DepHisListReq) (deployHistory []model.DeployHistory, err error) {
	db := common.DB.Model(&model.DeployHistory{})
	jobName := req.JobName
	if jobName != "" {
		db = db.Where("job_name = ?", fmt.Sprintf("%%%s%%", jobName))
	}
	createAtl := req.CreatedAtL
	if !createAtl.IsZero() {
		db = db.Where("created_at > ?", createAtl)
	}
	createAtr := req.CreatedAtR
	if !createAtr.IsZero() {
		db = db.Where("created_at < ?", createAtr)
	}
	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err = db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Find(&deployHistory).Error
	return deployHistory, err
}

// ListCount 获取符合条件的列表数量
func (receiver DeployHisService) ListCount(req *request.DepHisListReq) (count int64, err error) {
	db := common.DB.Model(&model.DeployHistory{})
	jobName := req.JobName
	if jobName != "" {
		db = db.Where("job_name = ?", fmt.Sprintf("%%%s%%", jobName))
	}
	createAtl := req.CreatedAtL
	if !createAtl.IsZero() {
		db = db.Where("created_at > ?", createAtl)
	}
	createAtr := req.CreatedAtR
	if !createAtr.IsZero() {
		db = db.Where("created_at < ?", createAtr)
	}
	err = db.Count(&count).Error
	return count, err
}

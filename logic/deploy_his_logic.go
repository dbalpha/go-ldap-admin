package logic

import (
	"fmt"
	"github.com/dbalpha/go-ldap-admin/model"
	"github.com/dbalpha/go-ldap-admin/model/request"
	"github.com/dbalpha/go-ldap-admin/model/response"
	"github.com/dbalpha/go-ldap-admin/public/tools"
	"github.com/dbalpha/go-ldap-admin/service/isql"
	"github.com/gin-gonic/gin"
)

/**
  @author: alpha
  @since: 2024/10/11
  @since: 2024/10/11
**/

type DeployHisLogic struct{}

func (receiver DeployHisLogic) Add(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {

	return nil, nil
}

// Delete 删除发布历史
func (receiver DeployHisLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.DeployDeleteReq)
	_ = c
	if !ok {
		return nil, ReqAssertErr
	}
	err := isql.DeployHis.Delete(r.DeployId)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("发布历史删除失败: %s", err.Error()))
	}
	return nil, nil
}

// List 获取发布历史记录
func (receiver DeployHisLogic) List(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.DepHisListReq)
	_ = c
	if !ok {
		return nil, ReqAssertErr
	}
	deployHiss, err := isql.DeployHis.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取发布历史失败: %s", err.Error()))
	}
	count, err := isql.DeployHis.ListCount(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取发布历史数量失败: %s", err.Error()))
	}
	return response.DeployHisListRsp{
		Total:      count,
		DeployHiss: deployHiss,
	}, nil
}

func (receiver DeployHisLogic) Info(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.DepHisInfo)
	_ = c
	if !ok {
		return nil, ReqAssertErr
	}
	deployhis := new(model.DeployHistory)
	filter := tools.H{"ID": r.DepHisId}
	err := isql.DeployHis.Find(filter, deployhis)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("查询发布历史信息错误: %s", err.Error()))
	}
	return deployhis, nil

}

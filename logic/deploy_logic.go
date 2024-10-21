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

type DeployLogic struct{}

// Add 添加job
func (receiver DeployLogic) Add(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.DeployAddReq)
	if !ok {
		return nil, ReqAssertErr
	}
	if isql.Deploy.Exist(tools.H{"Name": r.JobName}) {
		return nil, tools.NewMySqlError(fmt.Errorf("已存在同名job"))
	}
	_ = c
	deploy := model.Deploy{
		JobName:   r.JobName,
		GitUrl:    r.GitUrl,
		Branch:    r.Branch,
		RunConfig: r.RunConfig,
	}
	if deploy.Branch == "" {
		deploy.Branch = "master"
	}
	err := isql.Deploy.Add(&deploy)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("添加job失败"))
	}
	Jenkins.Add(&deploy)
	return nil, nil
}

// Delete 删除job
func (receiver DeployLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.DeployDeleteReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c
	err := isql.Deploy.Delete(r.DeployId)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("删除任务失败: %s", err.Error()))
	}
	return nil, nil
}

func (receiver DeployLogic) Update(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.DeployUpdateReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c
	filter := tools.H{"id": r.Id}
	oldDate := new(model.Deploy)
	err := isql.Deploy.Find(filter, oldDate)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("没有对应记录：%s", err.Error()))
	}
	deploy := model.Deploy{
		Model:     oldDate.Model,
		JobName:   r.JobName,
		GitUrl:    r.GitUrl,
		Branch:    r.Branch,
		RunConfig: r.RunConfig,
	}
	err = isql.Deploy.Update(&deploy)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("更新任务失败：%s", err.Error()))
	}
	Jenkins.Modify(&deploy)
	return nil, nil
}

func (receiver DeployLogic) List(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.DeployListReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c
	deploys, err := isql.Deploy.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取任务失败：%s", err.Error()))
	}
	rets := make([]model.Deploy, 0)
	for _, deploy := range deploys {
		rets = append(rets, *deploy)
	}
	count, err := isql.Deploy.ListCount(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取列表总数失败%s", err.Error()))
	}
	return response.DeployListRsp{
		Total:   count,
		Deploys: rets,
	}, nil
}

// Info 获取deploy信息
func (receiver DeployLogic) Info(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.DeployInfo)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c
	deploy := new(model.Deploy)
	filter := tools.H{"ID": r.DeployId}
	err := isql.Deploy.Find(filter, deploy)
	return deploy, err
}
func (receiver DeployLogic) Run(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.DeployRun)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c
	deployHis := model.DeployHistory{
		DeployID:   r.DeployId,
		DeployName: r.JobName,
	}
	err := isql.DeployHis.Add(&deployHis)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("发布历史写入失败：%s", err.Error()))
	}
	Jenkins.Run(r.JobName)
	return nil, nil
}

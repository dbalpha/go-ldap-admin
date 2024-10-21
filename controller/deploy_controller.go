package controller

import (
	"github.com/dbalpha/go-ldap-admin/logic"
	"github.com/dbalpha/go-ldap-admin/model/request"
	"github.com/gin-gonic/gin"
)

type DeployController struct{}

// Add 添加发布任务
// @Summary 添加发布任务
// @Description 添加发布任务
// @Tags 任务发布
// @Accept application/json
// @Produce application/json
// @Param  data body request.DeployAddReq true "添加用户记录的结构体"
// @Success 200 {object} response.ResponseBody
// @Router /deploy/add [post]
// @Security ApiKeyAuth
func (receiver DeployController) Add(c *gin.Context) {
	req := new(request.DeployAddReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Deploy.Add(c, req)
	})
}

// Delete 删除发布任务
// @Summary 删除发布任务
// @Description 删除发布任务
// @Tags 任务发布
// @Accept application/json
// @Produce application/json
// @Param  data body request.DeployDeleteReq true "添加用户记录的结构体"
// @Success 200 {object} response.ResponseBody
// @Router /deploy/delete [post]
// @Security ApiKeyAuth
func (receiver DeployController) Delete(c *gin.Context) {
	req := new(request.DeployDeleteReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Deploy.Delete(c, req)
	})
}

// Update 更新发布任务
// @Summary 更新发布任务
// @Description 更新发布任务
// @Tags 任务发布
// @Accept application/json
// @Produce application/json
// @Param  data body request.DeployUpdateReq true "添加用户记录的结构体"
// @Success 200 {object} response.ResponseBody
// @Router /deploy/update [post]
// @Security ApiKeyAuth
func (receiver DeployController) Update(c *gin.Context) {
	req := new(request.DeployUpdateReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Deploy.Update(c, req)
	})
}

// List 获取任务列表
// @Summary 任务列表
// @Description 获取任务列表
// @Tags 任务发布
// @Accept application/json
// @Produce application/json
// @Param  job_name query string false "job名字"
// @Success 200 {object} response.ResponseBody
// @Router /deploy/list [get]
// @Security ApiKeyAuth
func (receiver DeployController) List(c *gin.Context) {
	req := new(request.DeployListReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Deploy.List(c, req)
	})
}

// Info 获取任务详情
// @Summary 任务详情
// @Description 获取任务详情
// @Tags 任务发布
// @Tags 任务发布
// @Accept application/json
// @Produce application/json
// @Param  ID query string false "jobID"
// @Success 200 {object} response.ResponseBody
// @Router /deploy/info [get]
// @Security ApiKeyAuth
func (receiver DeployController) Info(c *gin.Context) {
	req := new(request.DeployInfo)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Deploy.Info(c, req)
	})
}

// DeployHisDelete 删除发布任务
// @Summary 删除发布任务
// @Description 删除发布任务
// @Tags 任务发布
// @Accept application/json
// @Produce application/json
// @Param  data body request.DeployDeleteReq true "添加用户记录的结构体"
// @Success 200 {object} response.ResponseBody
// @Router /deploy/history/delete [post]
// @Security ApiKeyAuth
func (receiver DeployController) DeployHisDelete(c *gin.Context) {
	req := new(request.DepHisDeleteReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.DeployHis.Delete(c, req)
	})
}

// DeployHisList 获取任务发布历史
// @Summary 任务发布历史
// @Description 获取任务发布历史
// @Tags 任务发布
// @Accept application/json
// @Produce application/json
// @Param  data body request.DeployInfo true "添加用户记录的结构体"
// @Success 200 {object} response.ResponseBody
// @Router /deploy/history/list [get]
// @Security ApiKeyAuth
func (receiver DeployController) DeployHisList(c *gin.Context) {
	req := new(request.DepHisListReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.DeployHis.List(c, req)
	})
}

// DeployHisInfo 获取任务详情
// @Summary 任务详情
// @Description 获取任务详情
// @Tags 任务发布
// @Tags 任务发布
// @Accept application/json
// @Produce application/json
// @Param  ID query string false "发布历史ID"
// @Success 200 {object} response.ResponseBody
// @Router /deploy/history/info [get]
// @Security ApiKeyAuth
func (receiver DeployController) DeployHisInfo(c *gin.Context) {
	req := new(request.DepHisInfo)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.DeployHis.Info(c, req)
	})
}

// Run  任务发布
// @Summary 任务发布
// @Description 任务发布
// @Tags 任务发布
// @Tags 任务发布
// @Accept application/json
// @Produce application/json
// @Param  ID query string false "发布历史ID"
// @Success 200 {object} response.ResponseBody
// @Router /deploy/run [get]
// @Security ApiKeyAuth
func (receiver DeployController) Run(c *gin.Context) {
	req := new(request.DeployRun)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Deploy.Run(c, req)
	})
}

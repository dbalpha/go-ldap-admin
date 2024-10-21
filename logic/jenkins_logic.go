package logic

import (
	"context"
	"fmt"
	"github.com/dbalpha/go-ldap-admin/model"
	"github.com/dbalpha/go-ldap-admin/public/common"
)

/**
  @author: alpha
  @since: 2024/10/12
**/

type JenkinsLogic struct{}

func (receiver JenkinsLogic) Add(deploy *model.Deploy) {
	ctx := context.Background()
	common.JenkinsCtx.CreateJob(ctx, "。。。。。。。。")
	fmt.Printf("%s任务已添加", deploy.JobName)
}
func (receiver JenkinsLogic) Modify(deploy *model.Deploy) {
	fmt.Printf("%s任务已修改", deploy.JobName)
}
func (receiver JenkinsLogic) Remove(jobName string) {
	fmt.Printf("%s任务已删除", jobName)
}
func (receiver JenkinsLogic) Run(jobName string) {
	fmt.Printf("%s任务已发布", jobName)

}

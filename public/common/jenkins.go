package common

/**
  @author: alpha
  @since: 2024/10/20
**/

import (
	"context"
	"github.com/bndr/gojenkins"
	"github.com/dbalpha/go-ldap-admin/config"
)

var JenkinsCtx *gojenkins.Jenkins

func InitJenkins() {
	ctx := context.Background()
	j, err := gojenkins.CreateJenkins(nil, config.Conf.Jenkins.Host+"/"+config.Conf.Jenkins.Port, config.Conf.Jenkins.User, config.Conf.Jenkins.Passwd).Init(ctx)
	if err != nil {
		Log.Errorf("failed to connect job: %v", err)
		return
	}
	JenkinsCtx = j
}

package feishu

import (
	"github.com/chyroc/lark"
	"github.com/dbalpha/go-ldap-admin/config"
)

func InitFeiShuClient() *lark.Lark {
	return lark.New(lark.WithAppCredential(
		config.Conf.FeiShu.AppID,
		config.Conf.FeiShu.AppSecret,
	))
}

package response

import "github.com/dbalpha/go-ldap-admin/model"

/*
*
	@author: alpha
	@since: 2024/10/11
*
*/

type DeployHisListRsp struct {
	Total      int64                 `json:"total"`
	DeployHiss []model.DeployHistory `json:"deployHiss"`
}

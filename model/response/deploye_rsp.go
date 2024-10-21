package response

import "github.com/dbalpha/go-ldap-admin/model"

type DeployListRsp struct {
	Total   int64          `json:"total"`
	Deploys []model.Deploy `json:"deploys"`
}

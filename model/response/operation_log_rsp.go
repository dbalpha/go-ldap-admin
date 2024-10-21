package response

import "github.com/dbalpha/go-ldap-admin/model"

type LogListRsp struct {
	Total int64                `json:"total"`
	Logs  []model.OperationLog `json:"logs"`
}

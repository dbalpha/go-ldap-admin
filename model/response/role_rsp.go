package response

import "github.com/dbalpha/go-ldap-admin/model"

type RoleListRsp struct {
	Total int64        `json:"total"`
	Roles []model.Role `json:"roles"`
}

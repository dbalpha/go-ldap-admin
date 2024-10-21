package response

import "github.com/dbalpha/go-ldap-admin/model"

type UserListRsp struct {
	Total int          `json:"total"`
	Users []model.User `json:"users"`
}

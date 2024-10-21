package request

import "time"

/**
  @author: alpha
  @since: 2024/10/11
**/

type DepHisDeleteReq struct {
	DepHisId []uint `json:"DepHis_id" validate:"required"`
}
type DepHisListReq struct {
	JobName    string    `json:"jobName" form:"jobName"`
	CreatedAtL time.Time `json:"createdAtL" form:"createdAtL"`
	CreatedAtR time.Time `json:"createdAtR" form:"createdAtR"`
	PageNum    int       `json:"pageNum" form:"pageNum"`
	PageSize   int       `json:"pageSize" form:"pageSize"`
}
type DepHisInfo struct {
	DepHisId uint `json:"depHisId" validate:"required"`
}

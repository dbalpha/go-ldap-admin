package request

type DeployAddReq struct {
	JobName   string `json:"jobName" validate:"required"`
	GitUrl    string `json:"gitUrl" validate:"required"`
	Branch    string `json:"branch"`
	RunConfig string `json:"runConfig"`
}

// DeployUpdateReq 更新job 结构体
type DeployUpdateReq struct {
	Id        uint   `json:"id" validate:"required"`
	JobName   string `json:"jobName" validate:"required"`
	GitUrl    string `json:"gitUrl" validate:"required"`
	Branch    string `json:"branch"`
	RunConfig string `json:"runConfig"`
}
type DeployDeleteReq struct {
	DeployId []uint `json:"deployIds" validate:"required"`
}
type DeployListReq struct {
	JobName  string `json:"jobName" form:"jobName"`
	GitUrl   string `json:"gitUrl" form:"gitUrl"`
	PageNum  int    `json:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}
type DeployInfo struct {
	DeployId uint `json:"deployId" validate:"required"`
}
type DeployRun struct {
	DeployInfo
	JobName string `json:"jobName" form:"jobName"`
}

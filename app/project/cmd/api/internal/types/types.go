// Code generated by goctl. DO NOT EDIT.
package types

type Project struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	FilePathJson string `json:"filePathJson"`
}

type GetProjectsReq struct {
}

type GetProjectsResp struct {
	Projects []Project `json:"projects"`
}

type CreateProjectReq struct {
	ProjectName  string `json:"projectName"`
	FilePathJson string `json:"filePathJson"`
}

type CreateProjectResp struct {
	Id int64 `json:"id"`
}

type UpdateProjectReq struct {
	ProjectId    int64  `json:"projectId"`
	ProjectName  string `json:"projectName"`
	FilePathJson string `json:"filePathJson"`
}

type UpdateProjectResp struct {
}

type DeleteProjectReq struct {
	ProjectId int64 `form:"projectId"`
}

type DeleteProjectResp struct {
}

type RunProjectReq struct {
	ProjectId   int64  `json:"projectId"`
	OssBucket   string `json:"ossBucket"`
	OssObject   string `json:"ossObject"`
	HandlerName string `json:"handlerName"`
}

type RunProjectResp struct {
}

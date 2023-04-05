// Code generated by goctl. DO NOT EDIT.
package types

type SearchDependenciesReq struct {
	Name     string `form:"name"`
	Page     int32  `form:"page"`
	PageSize int32  `form:"pageSize"`
}

type SearchDependenciesResp struct {
	Packages []string `json:"packages"`
	Total    int32    `json:"total"`
}

type GetDependenciesReq struct {
	Typ        int8   `form:"typ"`
	ProjectId  string `form:"projectId"`
	FunctionId string `form:"functionId"`
}

type GetDependenciesResp struct {
	PackageJson string `json:"packageJson"`
}

type DownloadDependencyReq struct {
	Typ          int8   `json:"typ"`
	ProjectId    string `json:"projectId"`
	FunctionId   string `json:"functionId"`
	Requirements string `json:"requirements"`
}

type DownloadDependencyResp struct {
}

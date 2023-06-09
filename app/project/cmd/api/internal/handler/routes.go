// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	project "mathless-backend/app/project/cmd/api/internal/handler/project"
	"mathless-backend/app/project/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: project.GetProjectsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: project.CreateProjectHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/",
				Handler: project.UpdateProjectHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/",
				Handler: project.DeleteProjectHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/run",
				Handler: project.RunProjectHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v1/project"),
	)
}

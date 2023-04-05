// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	function "mathless-backend/app/library/cmd/api/internal/handler/function"
	library "mathless-backend/app/library/cmd/api/internal/handler/library"
	"mathless-backend/app/library/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: library.GetLibrariesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: library.CreateLibraryHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/",
				Handler: library.UpdateLibraryHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/",
				Handler: library.DeleteLibraryHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/apply",
				Handler: library.ApplyLibraryHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/granted",
				Handler: library.GrantedLibraryHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/denied",
				Handler: library.DeniedLibraryHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v1/library"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: function.GetFunctionsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: function.CreateFunctionHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/",
				Handler: function.UpdateFunctionHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/",
				Handler: function.DeleteFunctionHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v1/function"),
	)
}
